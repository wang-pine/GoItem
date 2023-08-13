package Mydatabase

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql" //导入包但不使用，init()
)
var dbUsers *sql.DB
// 这里用来对单个用户的分表进行维护
func InitUsersDatabase() (err error) {
	fmt.Println("正在初始化用户视频列表数据库...\n")
	dsn := "douyin:123456@tcp(127.0.0.1:2206)/douyin_users"
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
	dbUsers.SetMaxIdleConns(10) //设置数据库连接池的最大连接数
	return
}
//根据用户的id创建每个用户的分表
func MakeNewUserTable(id int64){
	sqlStr := "CREATE TABLE `" +strconv.FormatInt(id,10)+"`("+
	"user_id BIGINT(20) NOT NULL,"+
    "video_id BIGINT(20) NOT NULL,"+
    "PRIMARY KEY(user_id)"+
	")ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;"
	_,err := db.Exec(sqlStr)
	if err != nil {
		fmt.Println("make table error:%v\n",err)
		return
	}
}
//创建完用户分表之后对用户分表插入视频id
//这个表现为用户每次上传完一个视频之后，就把这个视频的id插入到与用户同名的数据表中
func InsertVideoIdToUserTable(userId int64,videoId int64){
	sqlStr := "INSERT INTO `"+strconv.FormatInt(userId,10)+"`(user_id,video_id)VALLUES("+strconv.FormatInt(userId,10)+","+strconv.FormatInt(videoId,10)+")"
	ret,err := dbUsers.Exec(sqlStr)
	if err != nil {
		fmt.Println("insert failed,err%v\n", err)
		return
	}
	id,err := ret.LastInsertId()
	if err != nil{
		fmt.Println("get failed,err:%v\n", err)
		return
	}
	fmt.Println("插入成功的id是", id)
}
//回传用户的视频列表
func GetUserUploadVideos(userId int64)(ret ){

}