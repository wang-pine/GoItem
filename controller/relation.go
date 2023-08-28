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

type UserListResponse struct {
	common.Response
	UserList []common.User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	toUserId, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	actionType := c.Query("action_type")
	ok, userId := service.SearchToken(token)
	if !ok {
		fmt.Println("Token not exist")
	}
	if actionType == "1" {
		//1表示对目标用户进行关注
		Mydatabase.InsertFollowIdToUserTable(toUserId, userId)
		Mydatabase.InsertFollowerIdToUserTable(userId, toUserId)
		//改库1.总库2.视频库中的用户信息
		//1
		userInfo1 := Mydatabase.QueryUserById(userId)
		userInfo1.FollowCount++
		userInfo2 := Mydatabase.QueryUserById(toUserId)
		userInfo2.FollowerCount++
		Mydatabase.UpdateUser(&userInfo1)
		Mydatabase.UpdateUser(&userInfo2)
		//2
		videoList := Mydatabase.QueryVideoByAuthorId(userId)
		var i int
		for i = 0; i < len(videoList); i++ {
			videoList[i].AuthorFollowCount++
			Mydatabase.UpdateVideoInfo(&videoList[i])
		}
		videoList2 := Mydatabase.QueryVideoByAuthorId(toUserId)
		var j int
		for j = 0; j < len(videoList2); j++ {
			videoList2[j].AuthorFollowerCount++
			Mydatabase.UpdateVideoInfo(&videoList2[j])
		}

		c.JSON(http.StatusOK, common.Response{StatusCode: 0, StatusMsg: "关注成功"})
	} else if actionType == "2" {
		//2表示对目标用户进行取关
		Mydatabase.DeleteFollow(toUserId, userId)
		Mydatabase.DeleteFollower(userId, toUserId)
		//改库1.总库2.视频库中的用户信息
		//1
		userInfo1 := Mydatabase.QueryUserById(userId)
		userInfo1.FollowCount--
		userInfo2 := Mydatabase.QueryUserById(toUserId)
		userInfo2.FollowerCount--
		Mydatabase.UpdateUser(&userInfo1)
		Mydatabase.UpdateUser(&userInfo2)
		//2
		videoList := Mydatabase.QueryVideoByAuthorId(userId)
		var i int
		for i = 0; i < len(videoList); i++ {
			videoList[i].AuthorFollowCount--
			Mydatabase.UpdateVideoInfo(&videoList[i])
		}
		videoList2 := Mydatabase.QueryVideoByAuthorId(toUserId)
		var j int
		for j = 0; j < len(videoList2); j++ {
			videoList2[j].AuthorFollowerCount--
			Mydatabase.UpdateVideoInfo(&videoList2[j])
		}

		c.JSON(http.StatusOK, common.Response{StatusCode: 0, StatusMsg: "取消关注成功"})
	} else {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "关注失败"})
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	token := c.Query("token")
	_, id := service.SearchToken(token)
	if id != userId {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "无token"})
		return
	}
	list, len := Mydatabase.GetUserFollowList(userId)
	var followList []common.User
	var i int
	for i = 0; i < len; i++ {
		userInfo := Mydatabase.QueryUserById(list[i])
		var user common.User
		service.ConvertUserInfoToUser(&userInfo, &user, list[i])
		followList = append(followList, user)
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		UserList: followList,
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	token := c.Query("token")
	_, id := service.SearchToken(token)
	if id != userId {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "无token"})
		return
	}
	list, len := Mydatabase.GetUserFollowersList(userId)
	var followerList []common.User
	var i int
	for i = 0; i < len; i++ {
		userInfo := Mydatabase.QueryUserById(list[i])
		var user common.User
		service.ConvertUserInfoToUser(&userInfo, &user, list[i])
		followerList = append(followerList, user)
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		UserList: followerList,
	})
}

// FriendList all users have same friend list
// 朋友是互相关注的两个人为朋友
func FriendList(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	token := c.Query("token")
	_, id := service.SearchToken(token)
	if id != userId {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "无token"})
		return
	}
	list, len := Mydatabase.GetUserFollowList(userId)
	var i int
	//当关注列表中的用户有关注本user，即表示为相互关注，把其加入到friendList的id表中
	var friendList []int64
	var length int = 0
	for i = 0; i < len; i++ {
		if Mydatabase.IsFollow(userId, list[i]) {
			friendList = append(friendList, list[i])
			length++
		}
	}
	//加入到friendlist中之后进行user信息的返回
	var friendUserList []common.User
	var userInfo common.Userinfo
	var j int
	for j = 0; j < length; j++ {
		userInfo = Mydatabase.QueryUserById(friendList[j])
		var user common.User
		service.ConvertUserInfoToUser(&userInfo, &user, friendList[j])
		friendUserList = append(friendUserList, user)
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		UserList: friendUserList,
	})
}
