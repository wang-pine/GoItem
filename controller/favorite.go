package controller

import (
	"Mydatabase"
	"common"
	"github.com/gin-gonic/gin"
	"net/http"
	"service"
	"strconv"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	service.FavoriteAction(c)
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	//1.判断当前token是否是合法用户
	token := c.Query("token")
	res_tag, userId := service.SearchToken(token)
	if res_tag == false {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "User doesn't exist",
		})
		return
	}
	uId := c.Query("user_id")
	rId, err := strconv.ParseInt(uId, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "UserId不合法！",
		})
		return
	}
	if rId != userId {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "user和token不匹配！",
		})
		return
	}
	//2.查询用户喜欢的所有视频
	videoId, size := Mydatabase.GetUserFavoriteVideoList(userId)
	if size == 0 {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "该用户尚未点赞视频！",
		})
		return
	}
	videos := make([]common.Video, size)
	video := common.Video{}
	for i := 0; i < size; i++ {
		videoInfo := Mydatabase.QueryVideoById(videoId[i])
		service.ConvertVideoInfoToVideo(&videoInfo, &video, userId)
		videos = append(videos, video)
	}

	c.JSON(http.StatusOK, VideoListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		VideoList: videos,
	})
	return
}
