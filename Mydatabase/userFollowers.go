package Mydatabase

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql" //导入包但不使用，init()
)

var dbFollowers *sql.DB

// 这里的follower指的是追随者，是某个用户被哪些人关注了
// 千万不要迷糊了
// 这里用来对单个用户的分表进行维护
func InitFollowersDatabase() (err error) {
	fmt.Printf("正在初始化用户视频列表数据库...\n")
	dsn := "douyin:123456@tcp(127.0.0.1:3306)/douyin_followers"
	dbFollowers, err = sql.Open("mysql", dsn)
	//open函数是不会检查用户名和密码的
	if err != nil {
		return
	}
	err = dbFollowers.Ping() //尝试对数据库进行链接
	if err != nil {
		return
	}
	fmt.Println("链接数据库成功")
	dbFollowers.SetMaxIdleConns(100) //设置数据库连接池的最大连接数
	return
}

// 根据用户的id创建每个用户的分表
func MakeNewFollowerTable(id int64) (err error) {
	InitFollowersDatabase()
	sqlStr := "CREATE TABLE `" + strconv.FormatInt(id, 10) + "`(" +
		"follower_id BIGINT(20) NOT NULL," +
		"user_id BIGINT(20) NOT NULL," +
		"PRIMARY KEY(follower_id)" +
		")ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;"
	_, err1 := dbFollowers.Exec(sqlStr)
	if err1 != nil {
		fmt.Printf("make table error:%v\n", err1)
		return err1
	}
	return
}

// 创建完用户分表之后对用户分表插入视频id
// 这个表现为用户每次上传完一个视频之后，就把这个视频的id插入到与用户同名的数据表中
func InsertFollowerIdToUserTable(followerId int64, userId int64) {
	InitFollowersDatabase()
	sqlStr := "INSERT INTO `" + strconv.FormatInt(userId, 10) + "`(follower_id,user_id)VALUES(" + strconv.FormatInt(followerId, 10) + "," + strconv.FormatInt(userId, 10) + ");"
	ret, err := dbFollowers.Exec(sqlStr)
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

// 查询该用户的人员表
// id表是为了快速的知道用户的关注者id，这样查起总表来可以更快
func GetUserFollowersList(userId int64) (ret []int64, arrayLen int) {
	InitFollowersDatabase()
	var UserFollowersList []int64
	sqlStr := "SELECT follower_id,user_id FROM `" + strconv.FormatInt(userId, 10) + "` WHERE follower_id > 0"
	rows, err := dbFollowers.Query(sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer func() {
		rows.Close() // 释放数据库连接
	}()
	var user_id int64
	var follower_id int64
	for rows.Next() {
		err := rows.Scan(&follower_id, &user_id)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("scan success ,user id =%v", user_id)
		fmt.Printf("follower id = %v\n", follower_id)
		if follower_id != 0 {
			UserFollowersList = append(UserFollowersList, follower_id)
		}
	}
	return UserFollowersList, len(UserFollowersList)
}

// 查询user2有没有关注user1
func IsFollow(user1 int64, user2 int64) bool {
	arr, len := GetUserFollowersList(user1)
	//获取user1的追随者的列表
	var i int
	for i = 0; i < len; i++ {
		if arr[i] == user2 {
			return true
		}
	}
	return false
}

// 删除follower
func DeleteFollower(deleteFollowerId int64, userId int64) {
	InitFollowersDatabase()
	sqlStr := "DELETE FROM `" + strconv.FormatInt(userId, 10) + "`WHERE follower_id =" + strconv.FormatInt(deleteFollowerId, 10) + ";"
	_, err := dbFollowers.Exec(sqlStr)
	if err != nil {
		fmt.Println("error", err)
	}
}
