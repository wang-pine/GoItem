package controller

import (
	"common"
	"service"

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

// // PublishList all users have same publish video list
// func PublishList(c *gin.Context) {
// 	service.PublishList(c)
// }
