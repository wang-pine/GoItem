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
