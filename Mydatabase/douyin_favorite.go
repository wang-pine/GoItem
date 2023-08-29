package Mydatabase

/*
********************
存储用户点赞的视频
********************
*/
import (
	"config"
	"database/sql"
	"fmt"
	"strconv"
)

var dbFavorite *sql.DB

// 初始化视频点赞表
func InitFavoriteDatabase() (err error) {
	fmt.Printf("正在初始化视频用户点赞列表数据库...\n")
	dsn := "douyin:123456@tcp(" + config.GetDBAddr() + ")/douyin_favorite"
	dbFavorite, err = sql.Open("mysql", dsn)
	//open函数是不会检查用户名和密码的
	if err != nil {
		return
	}
	err = dbFavorite.Ping() //尝试对数据库进行链接
	if err != nil {
		return
	}
	fmt.Println("链接数据库成功")
	dbFavorite.SetMaxIdleConns(100) //设置数据库连接池的最大连接数
	return
}

// 根据用户id创建用户的点赞分表
// 这里需要传入用户的id
func MakeNewFavoriteTable(userId int64) (err error) {
	err = InitFavoriteDatabase()
	if err != nil {
		return err
	}
	sqlStr := "CREATE TABLE `" + strconv.FormatInt(userId, 10) + "`(" +
		"favorite_video_id BIGINT(20) NOT NULL," +
		"user_id BIGINT(20) NOT NULL," +
		"is_delete int(1) NOT NULL DEFAULT 0" +
		")ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;"
	_, err = dbFavorite.Exec(sqlStr)
	if err != nil {
		fmt.Printf("make table error:%v\n", err)
		return err
	}
	dbFavorite.Close()
	return
}

// 这是执行数据库语句的函数
// 用户不要调用
func execFavoriteDatabase(sqlStr string) bool {
	ret, err := dbFavorite.Exec(sqlStr)
	if err != nil {
		fmt.Printf("failed,err%v\n", err)
		return false
	}
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get failed,err:%v\n", err)
		return false
	}
	fmt.Println("运行成功的id是", id)
	return true
}

// 创建完用户分表之后对用户分表插入视频id
// 这个表现为用户每次上传完一个视频之后，就把这个视频的id插入到与用户同名的数据表中
func InsertUserIdToFavoriteTable(videoId int64, userId int64) bool {
	err := InitFavoriteDatabase()
	if err != nil {
		return false
	}
	videos, size := GetFavoriteVideoList(userId, videoId)
	if size == 0 {
		sqlStr := "INSERT INTO `" + strconv.FormatInt(userId, 10) + "`(favorite_video_id,user_id)VALUES(" + strconv.FormatInt(videoId, 10) + "," + strconv.FormatInt(userId, 10) + ");"
		return execFavoriteDatabase(sqlStr)
	}
	for i := 0; i < size; i++ {
		sqlStr := "UPDATE `" + strconv.FormatInt(userId, 10) + "` SET is_delete = 0" + " WHERE favorite_video_id = " + strconv.FormatInt(videos[i], 10)
		res := execFavoriteDatabase(sqlStr)
		if res == false {
			return false
		}
	}
	dbFavorite.Close()
	return true
}

// 逻辑删除，当delete表示1的时候表示删除
func DeleteUserIdToFavoriteTable(videoId int64, userId int64) bool {
	err := InitFavoriteDatabase()
	if err != nil {
		return false
	}
	videos, size := GetFavoriteVideoList(userId, videoId)
	if size == 0 {
		return false
	}
	for i := 0; i < size; i++ {
		sqlStr := "UPDATE `" + strconv.FormatInt(userId, 10) + "` SET is_delete = 1" + " WHERE favorite_video_id = " + strconv.FormatInt(videos[i], 10)
		res := execFavoriteDatabase(sqlStr)
		if res == false {
			return false
		}
	}
	dbFavorite.Close()
	return true
}

// 查询用户的喜欢的视频id表
func GetFavoriteVideoList(userId int64, videoId int64) (ret []int64, arrayLen int) {
	err := InitFavoriteDatabase()
	if err != nil {
		return nil, 0
	}
	var VideoFavorUsersList []int64
	sqlStr := "SELECT favorite_video_id,user_id,is_delete FROM `" + strconv.FormatInt(userId, 10) + "` WHERE favorite_video_id = " + strconv.FormatInt(videoId, 10)
	rows, err := dbFavorite.Query(sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			return
		} // 释放数据库连接
	}()
	var user_id int64
	var video_id int64
	var is_delete int64
	for rows.Next() {
		err := rows.Scan(&video_id, &user_id, &is_delete)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("scan success ,video id =%v", video_id)
		fmt.Printf("user id = %v\n", user_id)
		VideoFavorUsersList = append(VideoFavorUsersList, video_id)

	}
	dbFavorite.Close()
	return VideoFavorUsersList, len(VideoFavorUsersList)
}

// 查询用户的喜欢的视频id表
func GetUserFavoriteVideoList(userId int64) (ret []int64, arrayLen int) {
	err := InitFavoriteDatabase()
	if err != nil {
		return nil, 0
	}
	var VideoFavorUsersList []int64
	sqlStr := "SELECT favorite_video_id,user_id,is_delete FROM `" + strconv.FormatInt(userId, 10) + "` WHERE favorite_video_id > 0 "
	rows, err := dbFavorite.Query(sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			return
		} // 释放数据库连接
	}()
	var user_id int64
	var video_id int64
	var is_delete int64
	for rows.Next() {
		err := rows.Scan(&video_id, &user_id, &is_delete)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("scan success ,video id =%v", video_id)
		fmt.Printf("user id = %v\n", user_id)
		if is_delete == 0 {
			VideoFavorUsersList = append(VideoFavorUsersList, video_id)
		}

	}
	dbFavorite.Close()
	return VideoFavorUsersList, len(VideoFavorUsersList)
}
