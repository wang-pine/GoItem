package Mydatabase

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //导入包但不使用，init()
)

var dbPWD *sql.DB

// 链接账号密码数据表
func InitPWDDatabase() (err error) {
	fmt.Println("正在初始化用户账号密码数据库...\n")
	dsn := "douyin:123456@tcp(127.0.0.1:3306)/douyin_info"
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

// 插入注册用户的账号密码
func InsertNewUser(PWD string) (err error, id int64) {
	InitPWDDatabase()

}
//用户注册后账号密码需要进行加密处理
func StringToHash(PWD string)(Hash string){
	
}
// 数据库中查询用户的密码
func QueryUserPWD(id int64) (PWD string) {
	sqlStr := "select PWD FROM id_pwd WHERE id = ?"
	dbPWD.QueryRow(sqlStr, id).Scan(&PWD)
	return PWD
}

// 链接后端递送的数据，比对账号和密码
// 并判断是否正确
func JudgePWD(id int64, PWD string) (res bool) {

}
