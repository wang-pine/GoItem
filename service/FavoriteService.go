package service

import (
	"Mydatabase"
	"common"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	//1.判断当前token是否是合法用户
	token := c.Query("token")
	video_id := c.Query("video_id")
	res_tag, userId := SearchToken(token)
	fmt.Println(token)
	fmt.Println(video_id)
	if res_tag == false {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "User doesn't exist",
		})
		return
	}

	//2.判断当前操作
	actionType := c.Query("action_type")
	videoID, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 1,
			StatusMsg:  "videoId 不合法！",
		})
		return
	}
	//1是点赞操作，增加
	if actionType == "1" {
		res := Mydatabase.InsertUserIdToFavoriteTable(videoID, userId)
		//维护两个表
		Mydatabase.InsertUserIdToFavoriteTable(videoID, userId)
		Mydatabase.InsertUserIdToVideoTable(videoID, userId)
		//修改总表
		videoInfo := Mydatabase.QueryVideoById(videoID)
		videoInfo.VideoFavoriteCount++
		Mydatabase.UpdateVideoInfo(&videoInfo)

		if res == false {
			c.JSON(http.StatusOK, common.Response{
				StatusCode: 1,
				StatusMsg:  "插入数据库失败！",
			})
			fmt.Println("3")
			return
		}
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 0,
			StatusMsg:  "点赞成功！",
		})
		fmt.Println("4")
		return
	}

	if actionType == "2" {
		res := Mydatabase.DeleteUserIdToFavoriteTable(videoID, userId)
		Mydatabase.DeleteUserIdToVideoTable(videoID, userId)

		//修改总表
		videoInfo := Mydatabase.QueryVideoById(videoID)
		videoInfo.VideoFavoriteCount--
		Mydatabase.UpdateVideoInfo(&videoInfo)
		if res == false {
			c.JSON(http.StatusOK, common.Response{
				StatusCode: 1,
				StatusMsg:  "删除数据库失败！",
			})
			fmt.Println("5")
			return
		}
		c.JSON(http.StatusOK, common.Response{
			StatusCode: 0,
			StatusMsg:  "取消点赞成功！",
		})
		fmt.Println("6")
		return
	}

}

// FavoriteList all users have same favorite video list
//func FavoriteList(c *gin.Context) {
//	//1.判断当前token是否是合法用户
//	token := c.PostForm("token")
//	res_tag, userId := SearchToken(token)
//	if res_tag == false {
//		c.JSON(http.StatusOK, common.Response{
//			StatusCode: 1,
//			StatusMsg:  "User doesn't exist",
//		})
//		return
//	}
//	uId := c.PostForm("user_id")
//	rId, err := strconv.ParseInt(uId, 10, 64)
//	if err != nil {
//		c.JSON(http.StatusOK, common.Response{
//			StatusCode: 1,
//			StatusMsg:  "UserId不合法！",
//		})
//		return
//	}
//	if rId != userId {
//		c.JSON(http.StatusOK, common.Response{
//			StatusCode: 1,
//			StatusMsg:  "user和token不匹配！",
//		})
//		return
//	}
//	//2.查询用户喜欢的所有视频
//	videoId, size := Mydatabase.GetUserVideosList(userId)
//	if size == 0 {
//		c.JSON(http.StatusOK, common.Response{
//			StatusCode: 1,
//			StatusMsg:  "该用户尚未点赞视频！",
//		})
//		return
//	}
//	videos := make([]common.Video, size)
//	video := common.Video{}
//	for i := 0; i < size; i++ {
//		videoInfo := Mydatabase.QueryVideoById(videoId[i])
//		ConvertVideoInfoToVideo(&videoInfo, &video, userId)
//		videos = append(videos, video)
//	}
//
//	c.JSON(http.StatusOK, controller.VideoListResponse{
//		Response: common.Response{
//			StatusCode: 0,
//		},
//		VideoList: videos,
//	})
//	return
//}
