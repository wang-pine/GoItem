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

type VideoListResponse struct {
	common.Response
	VideoList []common.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory

func Publish(c *gin.Context) {
	service.Publish(c)
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	//token := c.Query("token")
	userId := c.Query("user_id")
	userID, _ := strconv.ParseInt(userId, 10, 64)
	//ok, userId := service.SearchToken(token)
	if userID != 0 {
		userVideosIdList, len := Mydatabase.GetUserVideosList(userID)
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
	} else {
		fmt.Println("不存在的id", userId)
		c.JSON(http.StatusOK, VideoListResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  "此人id不存在",
			},
		})
	}

}
