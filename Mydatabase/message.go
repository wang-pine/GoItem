package Mydatabase

import (
	"common"
	"config"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql" //导入包但不使用，init()
)

var dbMessage *sql.DB

// 本文件用于维护message数据库
func InitMessageDatabase() (err error) {
	fmt.Println("正在初始化信息数据库...")
	dsn := "douyin:123456@tcp(" + config.GetDBAddr() + ")/douyin_message"
	dbMessage, err = sql.Open("mysql", dsn)
	//open函数是不会检查用户名和密码的
	if err != nil {
		return
	}
	err = dbMessage.Ping() //尝试对数据库进行链接
	if err != nil {
		return
	}
	fmt.Println("链接数据库成功")
	dbMessage.SetMaxIdleConns(100) //设置数据库连接池的最大连接数
	return
}

// 创建用户分表
func MakeNewMessageTable(id int64) (err error) {
	InitMessageDatabase()
	sqlStr := "CREATE TABLE `" + strconv.FormatInt(id, 10) + "` (" +
		"message_id BIGINT(20) NOT NULL AUTO_INCREMENT," +
		"to_user_id BIGINT(20) NOT NULL," +
		"from_user_id BIGINT(20) NOT NULL," +
		"content VARCHAR(120) NOT NULL," +
		"create_time VARCHAR(60)NOT NULL," +
		"PRIMARY KEY(message_id)" +
		")ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;"
	_, err = dbMessage.Exec(sqlStr)
	if err != nil {
		return err
	}
	return
}
func InsertMessage(fromUserId int64, toUserId int64, content string) (date string, err error) {
	InitMessageDatabase()
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	sqlStr1 := "INSERT INTO `" +
		strconv.FormatInt(fromUserId, 10) +
		"`(to_user_id,from_user_id,content,create_time)VALUES(" +
		strconv.FormatInt(toUserId, 10) + "," +
		strconv.FormatInt(fromUserId, 10) + ",'" +
		content + "','" +
		currentTime + "');"
	sqlStr2 := "INSERT INTO `" +
		strconv.FormatInt(toUserId, 10) +
		"`(to_user_id,from_user_id,content,create_time)VALUES(" +
		strconv.FormatInt(toUserId, 10) + "," +
		strconv.FormatInt(fromUserId, 10) + ",'" +
		content + "','" +
		currentTime + "');"
	_, err = dbMessage.Exec(sqlStr1)
	if err != nil {
		return currentTime, err
	}
	_, err = dbMessage.Exec(sqlStr2)
	if err != nil {
		return currentTime, err
	}
	return
}

// 发送某个用户的全部消息列表
func GetMessageList(userId int64, msg_time int64) (messageList []common.Message) {
	InitMessageDatabase()
	sqlStr := "SELECT message_id,to_user_id,from_user_id,content,create_time FROM`" + strconv.FormatInt(userId, 10) + "`WHERE message_id >0 and create_time >" + strconv.FormatInt(msg_time, 10)
	rows, err := dbMessage.Query(sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer func() {
		rows.Close() // 释放数据库连接
	}()
	var messageId int64
	var toUserId int64
	var fromUserId int64
	var content string
	var createTime int64
	var message common.Message
	for rows.Next() {
		err := rows.Scan(&messageId, &toUserId, &fromUserId, &content, &createTime)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		message.Id = messageId
		message.ToUserId = toUserId
		message.FromUserId = fromUserId
		message.Content = content
		message.CreateTime = createTime
		messageList = append(messageList, message)
	}
	return messageList
}
