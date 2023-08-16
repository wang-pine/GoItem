package controller

import (
	"Mydatabase"
	"common"
	"fmt"
	"net/http"
	"path/filepath"
	"service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")

	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	user := usersLoginInfo[token]
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, common.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	//token := c.Query("token")
	userId := c.Query("user_id")
	userID, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		fmt.Println("convert userId error")
	}
	userVideosIdList, len := Mydatabase.GetUserVideosList(1)
	var i int
	var userVideoListDetailed []common.Video
	for i = 0; i < len; i++ {
		var temp common.Video
		videoInfoTemp := Mydatabase.QueryVideoById(userVideosIdList[i])
		service.ConvertVideoInfoToVideo(&videoInfoTemp, &temp, userID)
		userVideoListDetailed = append(userVideoListDetailed, temp)
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		VideoList: userVideoListDetailed,
	})
}
