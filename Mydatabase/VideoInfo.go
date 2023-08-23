package Mydatabase

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql" //导入包但不使用，init()
)

var dbVideos *sql.DB

// 这里用来对单个视频的分表进行维护
func InitVideosDatabase() (err error) {
	fmt.Printf("正在初始化视频用户点赞列表数据库...\n")
	dsn := "douyin:123456@tcp(127.0.0.1:3306)/douyin_videos"
	dbVideos, err = sql.Open("mysql", dsn)
	//open函数是不会检查用户名和密码的
	if err != nil {
		return
	}
	err = dbVideos.Ping() //尝试对数据库进行链接
	if err != nil {
		return
	}
	fmt.Println("链接数据库成功")
	dbVideos.SetMaxIdleConns(100) //设置数据库连接池的最大连接数
	return
}

// 根据视频的id创建每个视频的用户点赞的分表
// 这里需要传入视频的id
func MakeNewVideoTable(id int64) (err error) {
	InitVideosDatabase()
	sqlStr := "CREATE TABLE `" + strconv.FormatInt(id, 10) + "`(" +
		"user_id BIGINT(20) NOT NULL," +
		"video_id BIGINT(20) NOT NULL," +
		"is_delete int(1) NOT NULL default 0" +
		")ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;"
	_, err1 := dbVideos.Exec(sqlStr)
	if err1 != nil {
		fmt.Printf("make table error:%v\n", err)
		return err1
	}
	return
}

// 创建完用户分表之后对用户分表插入视频id
// 这个表现为用户每次上传完一个视频之后，就把这个视频的id插入到与用户同名的数据表中
func InsertUserIdToVideoTable(videoId int64, userId int64) {
	InitVideosDatabase()
	sqlStr := "INSERT INTO `" + strconv.FormatInt(videoId, 10) + "`(user_id,video_id)VALUES(" + strconv.FormatInt(userId, 10) + "," + strconv.FormatInt(videoId, 10) + ");"
	execVideoDatabase(sqlStr)
}

// 这是执行数据库语句的函数
// 用户不要调用
func execVideoDatabase(sqlStr string) {
	ret, err := dbVideos.Exec(sqlStr)
	if err != nil {
		fmt.Printf("failed,err%v\n", err)
		return
	}
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get failed,err:%v\n", err)
		return
	}
	fmt.Println("运行成功的id是", id)
}

// 查询视频的点赞用户id表
func GetFavoriteUsersList(videoId int64) (ret []int64, arrayLen int) {
	InitVideosDatabase()
	var VideoFavorUsersList []int64
	sqlStr := "SELECT user_id,video_id FROM `" + strconv.FormatInt(videoId, 10) + "` WHERE user_id > 0"
	rows, err := dbVideos.Query(sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer func() {
		rows.Close() // 释放数据库连接
	}()
	var user_id int64
	var video_id int64
	for rows.Next() {
		err := rows.Scan(&user_id, &video_id)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("scan success ,video id =%v", video_id)
		fmt.Printf("user id = %v\n", user_id)
		VideoFavorUsersList = append(VideoFavorUsersList, user_id)
	}
	return VideoFavorUsersList, len(VideoFavorUsersList)
}

// 检查这个用户是否喜欢了视频
func IsFavorite(UserID int64, VideoID int64) bool {
	usersList, length := GetFavoriteUsersList(VideoID)
	var i int
	for i = 0; i < length; i++ {
		if usersList[i] == UserID {
			return true
		}
	}
	return false
}

// 逻辑删除，当delete表示1的时候表示删除
func DeleteUserIdToVideoTable(videoId int64, userId int64) {
	err := InitVideosDatabase()
	if err != nil {
		return
	}
	sqlStr := "UPDATE `" + strconv.FormatInt(videoId, 10) + "` SET is_delete = 1" + " WHERE user_id = " + strconv.FormatInt(userId, 10)
	execVideoDatabase(sqlStr)
}
