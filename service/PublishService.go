package service

import (
	"Mydatabase"
	"common"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"path/filepath"
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
	new_video.Id = int64(rand.Int())
	new_video.Author = user
	new_video.CommentCount = 0
	new_video.CoverUrl = ""
	new_video.FavoriteCount = userinfo.FavoriteCount
	new_video.IsFavorite = false
	new_video.PlayUrl = "http://localhost:8888/static/" + finalName
	var videoInfo common.Videoinfo
	ConvertUserVideoToVideoIfo(&userinfo, &new_video, &videoInfo)
	videoInfo.VideoTitle = title
	videoInfo.VideoTime = ""
	res := Mydatabase.InsertVideoInfo(&videoInfo)
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

func PublishList(c *gin.Context) {
	//controller.userVideoList,len:=Mydatabase.GetUserVideosList()
	//var i int
	//userVideoListTotal []Video
	//for i=0;i<len;i++{
	//	// userVideoListTotal=append(userVideoList,)
	//}
	//c.JSON(http.StatusOK, VideoListResponse{
	//	Response: Response{
	//		StatusCode: 0,
	//	},
	//	VideoList: DemoVideos,
	//})
}
