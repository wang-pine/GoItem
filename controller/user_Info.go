package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserDTOResponse struct {
	Response
	UserDTO UserDTO `json:"userDTO"`
}

// 接口规定要的
type UserDTO struct {
	ID               int64  // 用户id
	Name             string // 用户名称
	Follow_count     int64  // 关注总数
	Follower_count   int64  // 粉丝总数
	Is_follow        bool   // true-已关注，false-未关注
	Avatar           string //用户头像
	Background_image string //用户个人页顶部大图
	Signature        string //个人简介
	Total_favorited  int64  //获赞数量
	Work_count       int64  //作品数量
	Favorite_count   int64  //点赞数量
}

var testUser = UserDTO{
	ID:               1,
	Name:             "test",
	Follow_count:     156,
	Follower_count:   4623456,
	Is_follow:        true,
	Avatar:           "https://img0.baidu.com/it/u=489552572,2707768722&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=500",
	Background_image: "https://img2.baidu.com/it/u=3485858097,1025969522&fm=253&fmt=auto&app=138&f=JPEG?w=890&h=500",
	Signature:        "这只是一条数据",
	Total_favorited:  56455156,
	Work_count:       30,
	Favorite_count:   45621,
}

var userInfos = map[int]UserDTO{
	int(testUser.ID): testUser,
}

func UserInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		log.Fatalf(err.Error())
	}
	//把user换成获取到的数据就行了
	if user, exist := userInfos[id]; exist {
		c.JSON(http.StatusOK, UserDTOResponse{
			Response: Response{StatusCode: 0},
			UserDTO:  user,
		})
	} else {
		c.JSON(http.StatusOK, UserDTOResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}

}
