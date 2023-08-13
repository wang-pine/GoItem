package Mydatabase

//用来对基本信息进行维护

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //导入包但不使用，init()
)

var db *sql.DB

// db表示连接池对象
func InitInfoDatabase() (err error) {
	fmt.Println("正在初始化用户数据库...\n")
	//用户名:密码@tcp(ip:端口)/数据库名字
	dsn := "douyin:123456@tcp(127.0.0.1:3306)/douyin_info"
	//链接数据集
	db, err = sql.Open("mysql", dsn)
	//open函数是不会检查用户名和密码的
	if err != nil {
		return
	}
	err = db.Ping() //尝试对数据库进行链接
	if err != nil {
		return
	}
	fmt.Println("链接数据库成功")
	db.SetMaxIdleConns(10) //设置数据库连接池的最大连接数
	return
}

// 向douyin_info的ID_PWD中插入PWD，同时返回被插入对象的ID
func InsertPWD(PWD string) (ret int64) {
	sqlStr := "insert into ID_PWD(PWD) VALUES(" + PWD + ");"
	//fmt.Println(sqlStr)
	//ret = id
	ret1, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Println("insert failed,err%v\n", err)
		return
	}
	id, err := ret1.LastInsertId()
	if err != nil {
		fmt.Println("get failed,err:%v\n", err)
		return
	}
	fmt.Println("id", id)
	ret = id
	return ret
}

// 更新用户数据
// 用id来查表更新用户数据
// 这个是原始函数，参数列表非常的长，最好不要调用
func UpdateUserAllInfo(ID int64, new_name string, new_follow_count int64, new_follower_count int64, new_avator int64, new_bg_img string, new_signature string, new_favor int64, new_work_count int64, new_favor_count int64) {
	fmt.Println("update")

}
