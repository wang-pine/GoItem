package Mydatabase

import (
	"common"
	"database/sql"
	"fmt"
	"strconv"
	"time"
)

var dbComment *sql.DB

// 初始化每个视频的评论数据库
func InitCommentDatabase() (err error) {
	fmt.Printf("正在初始化视频评论列表数据库...\n")
	dsn := "douyin:123456@tcp(127.0.0.1:3306)/douyin_comment"
	dbComment, err = sql.Open("mysql", dsn)
	//open函数是不会检查用户名和密码的
	if err != nil {
		return
	}
	err = dbComment.Ping() //尝试对数据库进行链接
	if err != nil {
		return
	}
	fmt.Println("链接数据库成功")
	dbComment.SetMaxIdleConns(100) //设置数据库连接池的最大连接数
	return
}

// 创建对应视频id的全部评论表
func MakeCommentTable(videoId int64) (err error) {
	InitCommentDatabase()
	sqlStr := "CREATE TABLE `" + strconv.FormatInt(videoId, 10) + "`(" +
		"comment_id BIGINT(20) NOT NULL AUTO_INCREMENT," +
		"status BOOLEAN NOT NULL," +
		"user_id BIGINT(20) NOT NULL ," +
		"content VARCHAR(120) NOT NULL," +
		"date VARCHAR(30) NOT NULL," +
		"PRIMARY KEY(comment_id)" +
		")ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;"
	_, err1 := dbComment.Exec(sqlStr)
	if err1 != nil {
		fmt.Printf("make table error:%v\n", err)
		return err1
	}
	return
}

// 在对应的视频列表下插入相关的评论
func InsertComment(videoId int64, userId int64, comment string) (id int64, date string) {
	InitCommentDatabase()
	//时间
	// year := time.Now().Year()
	// month := time.Now().Month()
	// day := time.Now().Day()

	// hour := time.Now().Hour()
	// minute := time.Now().Minute()
	// second := time.Now().Second()
	// currentDate := time.Date(year, month, day, hour, minute, second)
	currentDate := time.Now().Format("2006-01-02 15:04:05")
	sqlStr := "INSERT INTO `" +
		strconv.FormatInt(videoId, 10) +
		"`(status,user_Id,content,date)VALUES(" +
		"1," +
		strconv.FormatInt(userId, 10) + ",'" +
		comment + "','" +
		currentDate + "');"
	ret, err := dbComment.Exec(sqlStr)
	if err != nil {
		fmt.Printf("failed,err%v\n", err)
		return
	}
	id, err = ret.LastInsertId()
	if err != nil {
		fmt.Printf("get failed,err:%v\n", err)
		return
	}
	fmt.Println("运行成功的id是", id)
	return int64(id), currentDate
}

// 删除评论,本质上就是把评论的删除id置0
func DeleteComment(videoId int64, commentId int64) (err error) {
	InitCommentDatabase()
	sqlStr := "UPDATE `" + strconv.FormatInt(videoId, 10) + "`SET status = 0 WHERE comment_id =" + strconv.FormatInt(commentId, 10) + ";"
	ret, err := dbComment.Exec(sqlStr)
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
	return
}

// 发送评论列表
// 使用的时候请记得按照id修改个人用户的总信息
func GetCommentList(videoId int64) (commentList []common.Comment) {
	InitCommentDatabase()
	sqlStr := "SELECT comment_id,status,user_id,content,date FROM`" + strconv.FormatInt(videoId, 10) + "`WHERE comment_id >0"
	rows, err := dbComment.Query(sqlStr)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer func() {
		rows.Close() // 释放数据库连接
	}()
	var comment_id int64
	var status = -1
	var user_id int64
	var content string
	var date string
	var comment common.Comment
	for rows.Next() {
		err := rows.Scan(&comment_id, &status, &user_id, &content, &date)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		if status == 1 {
			comment.Id = comment_id
			comment.User.Id = user_id
			comment.Content = content
			comment.CreateDate = date
			commentList = append(commentList, comment)
		}
	}
	// var i int
	// for i = 0; i < len(commentList); i++ {
	// 	fmt.Println(commentList[i])
	// }
	return commentList
}
