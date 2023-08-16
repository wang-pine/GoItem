package controller

import (
	"Mydatabase"
	"common"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin

var usersLoginInfo = map[string]common.User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

/*
var userIdSequence = int64(1)
*/
type UserLoginResponse struct {
	common.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	common.Response
	User common.User `json:"user"`
}

/*
	func Register(c *gin.Context) {
		username := c.Query("username")
		password := c.Query("password")

		token := username + password

		if _, exist := usersLoginInfo[token]; exist {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: common.Response{StatusCode: 1, StatusMsg: "User already exist"},
			})
		} else {
			atomic.AddInt64(&userIdSequence, 1)
			newUser := common.User{
				Id:   userIdSequence,
				Name: username,
			}
			usersLoginInfo[token] = newUser
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: common.Response{StatusCode: 0},
				UserId:   userIdSequence,
				Token:    username + password,
			})
		}
	}

	func Login(c *gin.Context) {
		username := c.Query("username")
		password := c.Query("password")

		token := username + password

		if user, exist := usersLoginInfo[token]; exist {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: common.Response{StatusCode: 0},
				UserId:   user.Id,
				Token:    token,
			})
		} else {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
			})
		}
	}
*/
// MD5加密
func StringToMD5(PWD string) string {
	w := md5.New()
	w.Write([]byte(PWD))
	return hex.EncodeToString(w.Sum(nil))
}

// 用户注册函数
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	user := Mydatabase.QueryUserByName(username)
	if user.Id != 0 {
		fmt.Println("用户重名")
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: common.Response{StatusCode: 1, StatusMsg: "User Name already exist"},
		})
	}else{
		err, userId := Mydatabase.InsertNewUser(password)
		if err != nil {
			fmt.Println("插入新用户密码错误", err)
		}
		token := strconv.FormatInt(userId, 10) + StringToMD5(password)
		var userInfo common.Userinfo
		userInfo.Id = userId
		userInfo.Name = username
		res := Mydatabase.InsertUser(&userInfo)
		if !res {
			fmt.Println("插入用户总信息错误")
		}
		if service.PushToken(token, userId) != true {
			fmt.Println("insert token error")
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: common.Response{StatusCode: 1, StatusMsg: "User already exist"},
			})
		}
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: common.Response{StatusCode: 0, StatusMsg: "注册成功"},
			UserId:   userId,
			Token:    token,
		})
	}
}
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	userInfo := Mydatabase.QueryUserByName(username)
	if userInfo.Id == 0 {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: common.Response{
				StatusCode: 1, 
				StatusMsg: "Userxxxx doesn't exist",
			},
		})
		fmt.Println("用户不存在")
	} else {
		if StringToMD5(password) == Mydatabase.QueryUserPWD(userInfo.Id) {
			_, token := service.SearchTokenById(userInfo.Id)
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: common.Response{StatusCode: 0,StatusMsg: "密码正确，登录成功"},
				UserId:   userInfo.Id,
				Token:    token,
			})
			fmt.Println("密码正确能返回正确的值")
		}
		fmt.Println("用户存在")
	}
}
