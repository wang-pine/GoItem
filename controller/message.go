package controller

import (
	"Mydatabase"
	"common"
	"fmt"
	"net/http"
	"service"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var tempChat = map[string][]common.Message{}

var messageIdSequence = int64(1)

type ChatResponse struct {
	common.Response
	MessageList []common.MessageRender `json:"message_list"`
}

// MessageAction no practical effect, just check if token is valid
func MessageAction(c *gin.Context) {
	token := c.Query("token")
	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	content := c.Query("content")
	ok, userId := service.SearchToken(token)
	if !ok {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "token not exist"})
		return
	}
	_, err := Mydatabase.InsertMessage(userId, toUserId, content)
	if err != nil {
		fmt.Println("insert message error", err)
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "insert message error"})
		return
	}
	c.JSON(http.StatusOK, common.Response{StatusCode: 0, StatusMsg: "消息发送成功"})

	// if user, exist := usersLoginInfo[token]; exist {
	// 	userIdB, _ := strconv.Atoi(toUserId)
	// 	chatKey := genChatKey(user.Id, int64(userIdB))

	// 	atomic.AddInt64(&messageIdSequence, 1)
	// 	curMessage := common.Message{
	// 		Id:         messageIdSequence,
	// 		Content:    content,
	// 		CreateTime: time.Now().Format(time.Kitchen),
	// 	}

	// 	if messages, exist := tempChat[chatKey]; exist {
	// 		tempChat[chatKey] = append(messages, curMessage)
	// 	} else {
	// 		tempChat[chatKey] = []common.Message{curMessage}
	// 	}
	// 	c.JSON(http.StatusOK, common.Response{StatusCode: 0})
	// } else {
	// 	c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	// }
}

// userHistory用于存储当前用户发送之后最新的消息id
// 用以防止重复发送消息
var userHistory map[string]int = make(map[string]int)

//key:userId
//string是userId_user1_user2
//表示向当前user发送的1和2的全部聊天记录
/*
func MessageChat(c *gin.Context) {
	token := c.Query("token")
	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	ok, userId := service.SearchToken(token)
	c.JSON(http.StatusOK, common.Response{StatusCode: 0, StatusMsg: "正在请求消息数据库"})
	fmt.Println("请求消息数据库中")
	if !ok {
		fmt.Println("token error")
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "token not exist"})
		return
	}
	key := GetChatKey(userId, userId, toUserId)
	fmt.Println("当前key对应的值", userHistory[key])
	if userHistory[key] == 0 {
		userHistory[key] = -1
	}
	fmt.Println("之后key对应的值", userHistory[key])
	messageList := Mydatabase.GetMessageList(userId)
	if len(messageList) == 0 {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "msg List empty"})
		fmt.Println("message list = 0")
		return
	}
	var i int
	var temp common.MessageRender
	var MessageToRender []common.MessageRender
	for i = 0; i < len(messageList); i++ {
		if userHistory[key] < i {
			userHistory[key]++
			fmt.Println("当前key的值是", userHistory[key])
			temp.Id = messageList[i].Id
			temp.FromUserId = messageList[i].FromUserId
			temp.ToUserId = messageList[i].ToUserId
			temp.Content = messageList[i].Content
			//fmt.Println("当前createTime原本的值是", messageList[i].CreateTime)
			//temp.CreateTime = GetTime(messageList[i].CreateTime)
			//fmt.Println("转换之后createtime变成了", temp.CreateTime)
			times, _ := time.Parse("2006-01-02 15:04:05", messageList[i].CreateTime)
			timeUnix := times.Unix()
			fmt.Println("时间戳的内容是", timeUnix)
			temp.CreateTime = int(timeUnix)
			fmt.Println("当前message的值是", temp)
			MessageToRender = append(MessageToRender, temp)
		}
	}
	if len(MessageToRender) == 0 {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "no new msg"})
		return
	}
	fmt.Println("发送消息")
	c.JSON(http.StatusOK,
		ChatResponse{Response: common.Response{StatusCode: 0},
			MessageList: MessageToRender,
		})
	fmt.Println(MessageToRender)
}
*/
// 获取chatkey
func GetChatKey(userId int64, fromUserId int64, toUserId int64) (key string) {
	if fromUserId > toUserId {
		key = strconv.FormatInt(userId, 10) + "_" + strconv.FormatInt(fromUserId, 10) + "_" + strconv.FormatInt(toUserId, 10)
		return key
	} else {
		key = strconv.FormatInt(userId, 10) + "_" + strconv.FormatInt(toUserId, 10) + "_" + strconv.FormatInt(fromUserId, 10)
		return key
	}
}

// // 获取当前时间
// func GetTime(time string) int {
// 	fmt.Println("传入的time是", time)
// 	time1 := strings.Trim(time, " ")
// 	fmt.Println("去空格之后", time1)
// 	time2 := strings.Trim(time1, "-")
// 	fmt.Println("去-之后", time2)
// 	time3 := strings.Trim(time2, ":")
// 	fmt.Println("去:之后", time3)
// 	timeInt, _ := strconv.Atoi(time3)
// 	fmt.Println("最后", timeInt)
// 	return timeInt
// }

// MessageChat all users have same follow list

func MessageChat(c *gin.Context) {
	token := c.Query("token")
	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	ok, userId := service.SearchToken(token)
	if !ok {
		fmt.Println("token error")
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "token not exist"})
		return
	}
	//注意，这样发送消息的致命缺陷就是消息会重复发送
	//所以需要用到消息队列
	messageList := Mydatabase.GetMessageList(userId)
	//var res []common.Message
	var i int
	var messageRenderList []common.MessageRender
	var temp common.MessageRender
	for i = 0; i < len(messageList); i++ {
		if messageList[i].ToUserId == toUserId {
			//strconv.Atoi(messageList[i].CreateTime)
			time := strings.Trim(messageList[i].CreateTime, " ")
			time = strings.Trim(time, "-")
			time = strings.Trim(time, ":")
			timeInt, _ := strconv.Atoi(time)
			temp.Id = messageList[i].Id
			temp.FromUserId = messageList[i].FromUserId
			temp.ToUserId = messageList[i].ToUserId
			temp.Content = messageList[i].Content
			temp.CreateTime = timeInt
			messageRenderList = append(messageRenderList, temp)
			//res = append(res, messageList[i])
			//fmt.Println("此时的消息内容是", messageList[i])
		}
	}
	c.JSON(http.StatusOK,
		ChatResponse{Response: common.Response{StatusCode: 0},
			MessageList: messageRenderList,
		})
	fmt.Println(messageRenderList)
	// if user, exist := usersLoginInfo[token]; exist {
	// 	userIdB, _ := strconv.Atoi(toUserId)
	// 	chatKey := genChatKey(user.Id, int64(userIdB))

	// 	c.JSON(http.StatusOK, ChatResponse{Response: common.Response{StatusCode: 0}, MessageList: tempChat[chatKey]})
	// } else {
	// 	c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	// }
}
func genChatKey(userIdA int64, userIdB int64) string {
	if userIdA > userIdB {
		return fmt.Sprintf("%d_%d", userIdB, userIdA)
	}
	return fmt.Sprintf("%d_%d", userIdA, userIdB)
}
