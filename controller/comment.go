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

type CommentListResponse struct {
	common.Response
	CommentList []common.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	common.Response
	Comment common.Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	ok, _ := service.SearchToken(token)
	if !ok {
		fmt.Println("无此token")
		c.JSON(http.StatusOK, CommentListResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "token不存在"},
		})
		return
	}
	_, userId := service.SearchToken(token)
	actionType := c.Query("action_type") //1-发布，2-删除
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if actionType == "1" {
		commentText := c.Query("comment_text")
		commentId, currentDate := Mydatabase.InsertComment(videoId, userId, commentText)
		userInfo := Mydatabase.QueryUserById(userId)
		//视频总库信息修改
		videoInfo := Mydatabase.QueryVideoById(videoId)
		videoInfo.VideoCommentCount++
		ok := Mydatabase.UpdateVideoInfo(&videoInfo)
		if !ok {
			fmt.Println("数据库修改失败")
			c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "数据库修改失败"})
			return
		}
		var user common.User
		service.ConvertUserInfoToUser(&userInfo, &user, userId)
		c.JSON(http.StatusOK, CommentActionResponse{
			Response: common.Response{StatusCode: 0, StatusMsg: "发送评论成功"},
			Comment: common.Comment{
				Id:         commentId,
				User:       user,
				Content:    commentText,
				CreateDate: currentDate,
			},
		})
	} else if actionType == "2" {
		commentId, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64) //2的时候用
		err := Mydatabase.DeleteComment(videoId, commentId)
		if err != nil {
			fmt.Println("delete error")
		}
		//视频总库信息修改
		videoInfo := Mydatabase.QueryVideoById(videoId)
		videoInfo.VideoCommentCount--
		ok := Mydatabase.UpdateVideoInfo(&videoInfo)
		if !ok {
			fmt.Println("数据库修改失败")
			c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "数据库修改失败"})
			return
		}
		c.JSON(http.StatusOK, common.Response{StatusCode: 0, StatusMsg: "删除成功"})
	} else {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "error"})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	token := c.Query("token")
	ok, _ := service.SearchToken(token)
	if !ok {
		fmt.Println("无此token")
		c.JSON(http.StatusOK, CommentListResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "token不存在"},
		})
		return
	}
	//_, userId := service.SearchToken(token)
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	commentList := Mydatabase.GetCommentList(videoId)
	length := len(commentList)
	var i int
	var userId int64
	var user common.User
	for i = 0; i < length; i++ {
		userId = commentList[i].User.Id
		userInfo := Mydatabase.QueryUserById(userId)
		service.ConvertUserInfoToUser(&userInfo, &user, userId)
		commentList[i].User = user
	}
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    common.Response{StatusCode: 0},
		CommentList: commentList,
	})
}
