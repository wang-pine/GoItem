package Msg

import (
	"common"
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
)
var MsgRouter = gin.Default()
var MsgList []common.Message
var mutex sync.Mutex
//生成msg密钥
func GenerateMegKey(fromUser int64, toUser int64, msgId int64) {

}
//启动消息队列服务（放入main函数中常驻）
func RunMessageServer() {
	err := MsgRouter.Run(":8080")
	if err != nil{
		fmt.Println("run message server error",err)
	}
	
	for {
		err := getMsg()
		if err != nil{
			fmt.Println("get Msg error%v\n",err)
			continue
		}
		go postMsg()
	}
}
//获得队首信息
func getMsg()(err error){
	router.GET("/message",func(c *gin.Context)){
		mutex.Lock()
		if len(msgList) == 0 {
			c.JSON(200, gin.H{"msg": "no message"})
			mutex.Unlock()
			return
		}
		//取出最早的信息
		msg:=MsgList[0]
		MsgList = MsgList[1:]
		mutex.Unlock()
		c.JSON(200, gin.H{"msg": msg})
	}
}
//向消息队列中加入信息
func postMsg(){
	router.POST("/message",func(c *gin.Context)){
		id,exists := c.GetPostForm("id")
		if !exists{
			return
		}
		fromUserId,exists := c.GetPostForm("from_user_id")
		if !exists{
			return
		}
		toUserId,exists := c.GetPostForm("to_user_id")
		if !exists{
			return
		}
		content,exists:=c.GetPostForm("content")
		if !exists{
			return
		}
		createTime,exists := c.GetPostForm("create_time")
		if !exists{
			return
		}
		var temp common.Message
		temp.Id =  id
		temp.FromUserId = fromUserId
		temp.ToUserId = toUserId
		temp.Content = comtent
		temp.CreateTime = createTime
		mutex.Lock()
		MsgList = append(MsgList, temp)
		mutex.Unlock()
		c.JSON(200,gin.H{"msg":"added"})
	}
}