package Msg

// import (
// 	"common"
// 	"fmt"
// 	"strconv"
// 	"sync"

// 	"github.com/gin-gonic/gin"
// )

// // 队列中的消息形式
// // 消息由唯一的key和消息组成，只有以key来请求消息队列才会发送最新的消息
// // 每次发完消息之后将用户id侧的信息进行更新，保证每个用户接收到的都是最新的信息
// type msgInfo struct {
// 	key string
// 	msg common.Message
// }

// var MsgRouter = gin.Default()

// // 消息队列
// var MsgList []msgInfo

// var mutex sync.Mutex

// // 用于存储最新的msg信息的id，id是以时间顺序存储的
// // 这样放入队列中的msg保持在发送后的最新状态
// // 存储的内容是用户id:最新的消息id
// // 消息内容在发送之后更新
// var msgHistory map[int64]int

// // 生成msg密钥
// func GenerateMsgKey(fromUser int64, toUser int64) (key string) {
// 	if fromUser > toUser {
// 		key = strconv.FormatInt(fromUser, 10) + "_" + strconv.FormatInt(toUser, 10)
// 	} else {
// 		key = strconv.FormatInt(toUser, 10) + "_" + strconv.FormatInt(fromUser, 10)
// 	}
// 	return key
// }

// // 启动消息队列服务（放入main函数中常驻）
// func RunMessageServer() {
// 	err := MsgRouter.Run(":8080")
// 	if err != nil {
// 		fmt.Println("run message server error", err)
// 	}

// 	for {
// 		err := getMsg()
// 		if err != nil {
// 			fmt.Println("get Msg error%v\n", err)
// 			continue
// 		}
// 		go postMsg()
// 	}
// }

// // 获得符合需求的信息
// func getMsg() (err error) {
// 	MsgRouter.GET("/message", func(c *gin.Context) {
// 		key := c.Query("key")
// 		userId,_ := strconv.ParseInt(c.Query("user_id"), 10, 64)
// 		msgId := msgHistory[userId]
// 		var i int
// 		mutex.Lock()
// 		if len(MsgList) == 0 {
// 			c.JSON(200, gin.H{"msg": "no message"})
// 			mutex.Unlock()
// 			return
// 		}
// 		//取出符合需求的信息
// 		msg := MsgList[0]
// 		MsgList = MsgList[1:]
// 		mutex.Unlock()
// 		c.JSON(200, gin.H{"msg": msg})
// 	})
// }

// // 向消息队列中加入信息
// func postMsg() {
// 	MsgRouter.POST("/message", func(c *gin.Context) {
// 		key, exists := c.GetPostForm("key")
// 		if !exists {
// 			return
// 		}
// 		userId, exists := c.GetPostForm("user_id")
// 		if !exists {
// 			return
// 		}
// 		toUserId, exists := c.GetPostForm("to_user_id")
// 		if !exists {
// 			return
// 		}
// 		content, exists := c.GetPostForm("content")
// 		if !exists {
// 			return
// 		}
// 		createTime, exists := c.GetPostForm("create_time")
// 		if !exists {
// 			return
// 		}
// 		var temp common.Message
// 		temp.Id, _ = strconv.ParseInt(id, 10, 64)
// 		temp.FromUserId, _ = strconv.ParseInt(fromUserId, 10, 64)
// 		temp.ToUserId, _ = strconv.ParseInt(toUserId, 10, 64)
// 		temp.Content = content
// 		temp.CreateTime = createTime
// 		mutex.Lock()
// 		temp1 := msgInfo{
// 			key: GenerateMsgKey(temp.FromUserId, temp.ToUserId),
// 			msg: temp,
// 		}
// 		MsgList = append(MsgList, temp1)
// 		mutex.Unlock()
// 		c.JSON(200, gin.H{"msg": "added"})
// 	})
// }

// // 如果检查到userid的history为0的时候，就往该用户的消息队列中增加东西
// // 否则就是直接在内存中向前台发信息（增加速度）
// // 然后存到数据库中即可
// func checkUserHistory() {

// }
