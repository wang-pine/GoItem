package service

import (
	"Mydatabase"
	"common"
	"config"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	common.Response
	VideoList []common.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")
	res_tag, userId := SearchToken(token)
	title := c.PostForm("title")
	fmt.Println("获取到的token是：" + token)
	if res_tag == false {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "User doesn't exist",
		})
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
	finalName := fmt.Sprintf("%d_%s", userId, filename)
	saveFile := filepath.Join("./public/", finalName)

	if title == "" {
		title = finalName
	}
	var userinfo common.Userinfo
	userinfo = Mydatabase.QueryUserById(userId)
	var user common.User
	ConvertUserInfoToUser(&userinfo, &user, userId)
	new_video := common.Video{}
	//获取最后一个视频的id
	last := Mydatabase.GetLastVideo()
	new_video.Id = last.VideoId + 1 //应当是最后一个+1
	new_video.Author = user
	new_video.CommentCount = 0
	new_video.CoverUrl = GetRandCover()
	new_video.FavoriteCount = userinfo.FavoriteCount
	new_video.IsFavorite = false
	new_video.PlayUrl = config.GetLocalAddr() + "/static/" + finalName
	//new_video.PlayUrl = "http://192.168.3.10:8888/static/" + finalName
	var videoInfo common.Videoinfo
	ConvertUserVideoToVideoIfo(&userinfo, &new_video, &videoInfo)
	videoInfo.VideoTitle = title
	videoInfo.VideoTime = ""
	//插入到视频总表的数据库
	videoInfo.AuthorWorkCount++
	res := Mydatabase.InsertVideoInfo(&videoInfo)
	//插入进用户分表
	Mydatabase.InsertVideoIdToUserTable(videoInfo.VideoId, videoInfo.AuthorId)
	//创建视频点赞表
	err = Mydatabase.MakeNewVideoTable(videoInfo.VideoId)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "创建视频点赞表失败！",
		})
		return
	}
	//创建视频评论表
	err = Mydatabase.MakeCommentTable(videoInfo.VideoId)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "创建视频评论表失败！",
		})
		return
	}
	if res == false {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "添加数据库有误！",
		})
		return
	}

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
