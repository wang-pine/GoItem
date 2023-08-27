package Mydatabase

import (
	"config"
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql" //导入包但不使用，init()
)

var dbFollow *sql.DB

// 维护用户关注的人
func InitFollowDatabase() (err error) {
	fmt.Printf("正在初始化用户视频列表数据库...\n")
	dsn := "douyin:123456@tcp(" + config.GetDBAddr() + ")/douyin_follow"
	dbFollow, err = sql.Open("mysql", dsn)
	//open函数是不会检查用户名和密码的
	if err != nil {
		return
	}
	err = dbFollow.Ping() //尝试对数据库进行链接
	if err != nil {
		return
	}
	fmt.Println("链接数据库成功")
	dbFollow.SetMaxIdleConns(100) //设置数据库连接池的最大连接数
	return
}

// 根据用户的id创建每个用户的分表
func MakeNewFollowTable(id int64) (err error) {
	InitFollowDatabase()
	sqlStr := "CREATE TABLE `" + strconv.FormatInt(id, 10) + "`(" +
		"follow_id BIGINT(20) NOT NULL," +
		"user_id BIGINT(20) NOT NULL," +
		"PRIMARY KEY(follow_id)" +
		")ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;"
	_, err1 := dbFollow.Exec(sqlStr)
	if err1 != nil {
		fmt.Printf("make table error:%v\n", err1)
		return err1
	}
	return
}

// 添加用户的关注
func InsertFollowIdToUserTable(followId int64, userId int64) {
	InitFollowDatabase()
	sqlStr := "INSERT INTO `" + strconv.FormatInt(userId, 10) + "`(follow_id,user_id)VALUES(" + strconv.FormatInt(followId, 10) + "," + strconv.FormatInt(userId, 10) + ");"
	ret, err := dbFollow.Exec(sqlStr)
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

// 查询该用户关注表
func GetUserFollowList(userId int64) (ret []int64, arrayLen int) {
	InitFollowDatabase()
	var UserFollowList []int64
	sqlStr := "SELECT follow_id,user_id FROM `" + strconv.FormatInt(userId, 10) + "` WHERE follow_id > 0"
	rows, err := dbFollow.Query(sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer func() {
		rows.Close() // 释放数据库连接
	}()
	var user_id int64
	var follow_id int64
	for rows.Next() {
		err := rows.Scan(&follow_id, &user_id)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("scan success ,user id =%v", user_id)
		fmt.Printf("follower id = %v\n", follow_id)
		if follow_id != 0 {
			UserFollowList = append(UserFollowList, follow_id)
		}
	}
	return UserFollowList, len(UserFollowList)
}

// 删除用户的关注者，本质上就是将用户关注者的id置0，扫描的时候会对id=0进行忽略
func DeleteFollow(deleteFollowId int64, userId int64) {
	InitFollowDatabase()
	sqlStr := "DELETE FROM `" + strconv.FormatInt(userId, 10) + "`WHERE follow_id =" + strconv.FormatInt(deleteFollowId, 10) + ";"
	_, err := dbFollow.Exec(sqlStr)
	if err != nil {
		fmt.Println("error", err)
	}
}
