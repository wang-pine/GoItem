package controller

import (
	"Mydatabase"
	"common"
	"net/http"
	"service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	common.Response
	VideoList []common.Video `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	token := c.Query("token")
	var VideoList []common.Video
	LatestVideo := Mydatabase.GetLastVideo()
	tailVideoId := LatestVideo.VideoId
	_, userId := service.SearchToken(token)
	var temp common.Video
	service.ConvertVideoInfoToVideo(&LatestVideo, &temp, userId)
	VideoList = append(VideoList, temp)
	var i int
	lastId, _ := strconv.Atoi(strconv.FormatInt(tailVideoId, 10))
	for i = lastId - 1; i > 0; i-- {
		videoTemp := Mydatabase.QueryVideoById(int64(i))
		var temp1 common.Video
		service.ConvertVideoInfoToVideo(&videoTemp, &temp1, userId)
		VideoList = append(VideoList, temp1)
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  common.Response{StatusCode: 0},
		VideoList: VideoList,
		NextTime:  time.Now().Unix(),
	})
}
