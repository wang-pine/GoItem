package Mydatabase

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //导入包但不使用，init()
)

var dbPWD *sql.DB

// 链接账号密码数据表
func InitPWDDatabase() (err error) {
	fmt.Printf("正在初始化用户账号密码数据库...\n")
	dsn := "douyin:123456@tcp(127.0.0.1:3306)/douyin_info"
	dbPWD, err = sql.Open("mysql", dsn)
	//注意这里是=，因为go会自动检查局部变量覆盖全局变量，切记
	//open函数是不会检查用户名和密码的
	if err != nil {
		return err
	}
	err = dbPWD.Ping() //尝试对数据库进行链接
	if err != nil {
		return err
	}
	fmt.Println("链接数据库成功")
	dbPWD.SetMaxIdleConns(10) //设置数据库连接池的最大连接数
	return err
}

// 插入注册用户的账号密码
//这里传入字符串就行，会自动加密保存到数据库中
func InsertNewUser(PWD string) (err error) {
	InitPWDDatabase()
	md5Str := StringToMD5(PWD)
	//fmt.Println(md5Str)
	sqlStr := "INSERT INTO id_pwd(PWD) VALUES('" + md5Str + "')"
	ret, err := dbPWD.Exec(sqlStr)
	if err != nil {
		fmt.Printf("insert failed,err%v\n", err)
		return
	}
	//id用来获取当前插入的序列id
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get failed,err:%v\n", err)
		return
	}
	fmt.Println("插入成功id=", id)
	return err
}

// 用户注册后账号密码需要进行加密处理
// MD5加密
func StringToMD5(PWD string) string {
	w := md5.New()
	w.Write([]byte(PWD))
	return hex.EncodeToString(w.Sum(nil))
}

// 数据库中查询用户的密码
func QueryUserPWD(id int64) (PWD string) {
	InitPWDDatabase()
	sqlStr := "select PWD FROM id_pwd WHERE id = ?"
	dbPWD.QueryRow(sqlStr, id).Scan(&PWD)
	return PWD
}

// 链接后端递送的数据，比对账号和密码
// 并判断是否正确
//这里传入用户id和想要进行比对的密码的字符串即可
func JudgePWD(id int64, PWD string) (res bool) {
	InitPWDDatabase()
	hash:=StringToMD5(PWD)
	passWord := QueryUserPWD(id)
	return hash == passWord
}
