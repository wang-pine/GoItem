package controller

import (
	"Mydatabase"
	"common"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
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

type UserDTOResponse struct {
	common.Response
	User common.User `json:"user"`
}

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
	} else {
		err, userId := Mydatabase.InsertNewUser(password)
		if err != nil {
			fmt.Println("插入新用户密码错误", err)
		}
		token := service.CreateUserToken(userId, password)
		//新建userInfo结构体
		//用于插入到信息总库中
		var userInfo common.Userinfo
		userInfo.Id = userId
		userInfo.Name = username
		userInfo.Avator = service.GetRandAvatar()
		userInfo.BackgroundImage = service.GetRandBGIMG()
		//在总数据库中插入当前用户的信息
		res := Mydatabase.InsertUser(&userInfo)
		if !res {
			fmt.Println("插入用户总信息错误")
		}
		//创建用户上传视频的分表
		err1 := Mydatabase.MakeNewUserTable(userId)
		//创建用户喜欢的分表
		err = Mydatabase.MakeNewFavoriteTable(userId)
		if err != nil {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: common.Response{StatusCode: 1, StatusMsg: "创建用户喜欢列表出错"},
			})
			return
		}
		if err1 != nil {
			fmt.Println("创建用户分表错误", err1)
		}
		err2 := Mydatabase.MakeNewFollowTable(userId)
		if err2 != nil {
			fmt.Println("创建用户关注分表错误", err2)
		}
		err3 := Mydatabase.MakeNewFollowerTable(userId)
		if err3 != nil {
			fmt.Println("创建用户被关注分表错误", err3)
		}
		err4 := Mydatabase.MakeNewMessageTable(userId)
		if err4 != nil {
			fmt.Println("创建用户信息发送分表错误", err4)
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
				StatusMsg:  "Userxxxx doesn't exist",
			},
		})
		fmt.Println("用户不存在")
	} else {
		if StringToMD5(password) == Mydatabase.QueryUserPWD(userInfo.Id) {
			ok, token := service.SearchTokenById(userInfo.Id)
			if !ok {
				token = service.CreateUserToken(userInfo.Id, Mydatabase.QueryUserPWD(userInfo.Id))

			}
			service.PushToken(token, userInfo.Id)
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: common.Response{StatusCode: 0,
					StatusMsg: "密码正确，登录成功",
				},
				UserId: userInfo.Id,
				Token:  token,
			})
			fmt.Println(token)
			fmt.Println("密码正确能返回正确的值")
		}
		fmt.Println("用户存在")
	}
}
func UserInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		log.Fatalf(err.Error())
	}
	//把user换成获取到的数据就行了
	userInfo := Mydatabase.QueryUserById(int64(id))
	var user common.User
	service.ConvertUserInfoToUser(&userInfo, &user, int64(id))
	if userInfo.Id != 0 {
		c.JSON(http.StatusOK, UserDTOResponse{
			Response: common.Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserDTOResponse{
			Response: common.Response{
				StatusCode: 1,
				StatusMsg:  "User doesn't exist",
			},
		})
	}

}
