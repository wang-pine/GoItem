package Mydatabase
/*
********************
存储用户投稿的视频
********************
*/
import (
	"config"
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql" //导入包但不使用，init()
)

var dbUsers *sql.DB

// 这里用来对单个用户的分表进行维护
//单个用户的分表存放的是该用户上传的视频
func InitUsersDatabase() (err error) {
	fmt.Printf("正在初始化用户视频列表数据库...\n")
	dsn := "douyin:123456@tcp(" + config.GetDBAddr() + ")/douyin_users"
	dbUsers, err = sql.Open("mysql", dsn)
	//open函数是不会检查用户名和密码的
	if err != nil {
		return
	}
	err = dbUsers.Ping() //尝试对数据库进行链接
	if err != nil {
		return
	}
	fmt.Println("链接数据库成功")
	dbUsers.SetMaxIdleConns(100)

	//设置数据库连接池的最大连接数
	return
}

// 根据用户的id创建每个用户的分表
func MakeNewUserTable(id int64) (err error) {
	InitUsersDatabase()
	sqlStr := "CREATE TABLE `" + strconv.FormatInt(id, 10) + "`(" +
		"video_id BIGINT(20) NOT NULL," +
		"user_id BIGINT(20) NOT NULL," +
		"PRIMARY KEY(video_id)" +
		")ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;"
	_, err1 := dbUsers.Exec(sqlStr)
	if err1 != nil {
		fmt.Printf("make table error:%v\n", err)
		return err1
	}
	return
}

// 创建完用户分表之后对用户分表插入视频id
// 这个表现为用户每次上传完一个视频之后，就把这个视频的id插入到与用户同名的数据表中
func InsertVideoIdToUserTable(videoId int64, userId int64) {
	InitUsersDatabase()
	sqlStr := "INSERT INTO `" + strconv.FormatInt(userId, 10) + "`(video_id,user_id)VALUES(" + strconv.FormatInt(videoId, 10) + "," + strconv.FormatInt(userId, 10) + ");"
	execDatabase(sqlStr)
}

// 这是执行数据库语句的函数
// 用户不要调用
func execDatabase(sqlStr string) {
	ret, err := dbUsers.Exec(sqlStr)
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

// 查询视频的id表
// 还需要向总库请求视频的具体列表
// id表是为了快速的知道用户的视频id，这样查起总表来可以更快
func GetUserVideosList(userId int64) (ret []int64, arrayLen int) {
	InitUsersDatabase()
	var UserVideoList []int64
	sqlStr := "SELECT video_id,user_id FROM `" + strconv.FormatInt(userId, 10) + "` WHERE video_id > 0"
	rows, err := dbUsers.Query(sqlStr)
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
		err := rows.Scan(&video_id, &user_id)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("scan success ,user id =%v", user_id)
		fmt.Printf("viideo id = %v\n", video_id)
		UserVideoList = append(UserVideoList, video_id)
	}
	return UserVideoList, len(UserVideoList)
}
