package controller

import (
	"Mydatabase"
	"common"
	"fmt"
	"net/http"
	"service"
	"strconv"

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
}

// userHistory用于存储当前用户发送之后最新的消息id
// 用以防止重复发送消息
//var userHistory map[string]int = make(map[string]int)

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


// MessageChat all users have same follow list

func MessageChat(c *gin.Context) {
	token := c.Query("token")
	msg_time, _ := strconv.ParseInt(c.Query("pre_msg_time"), 10, 64)
	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	ok, userId := service.SearchToken(token)
	// userId = 2
	if !ok {
		fmt.Println("token error")
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "token not exist"})
		return
	}
	//注意，这样发送消息的致命缺陷就是消息会重复发送
	//所以需要用到消息队列
	messageList := Mydatabase.GetMessageList(userId, msg_time)
	//var res []common.Message
	var i int
	var messageRenderList []common.MessageRender
	var temp common.MessageRender
	for i = 0; i < len(messageList); i++ {
		if messageList[i].ToUserId == toUserId {
			temp.FromUserId = messageList[i].FromUserId
			temp.ToUserId = messageList[i].ToUserId
		} else {
			temp.FromUserId = messageList[i].ToUserId
			temp.ToUserId = messageList[i].FromUserId
		}
		temp.Id = messageList[i].Id
		temp.Content = messageList[i].Content
		temp.CreateTime = messageList[i].CreateTime
		messageRenderList = append(messageRenderList, temp)
	}
	c.JSON(http.StatusOK,
		ChatResponse{Response: common.Response{StatusCode: 0},
			MessageList: messageRenderList,
		})
	fmt.Println(messageRenderList)
}
