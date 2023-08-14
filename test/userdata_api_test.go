package test

import (
	"Mydatabase"
	"fmt"
	"testing"
)

// 数据库连接测试
func TestGetDB(t *testing.T) {
	//e := newExpect(t)
	Mydatabase.GetDB()
}

// 根据id查询用户测试
func TestGetUserId(t *testing.T) {
	//e := newExpect(t)
	fmt.Println(Mydatabase.QueryUserById(20))
}

// 根据id查询用户测试
func TestGetUserName(t *testing.T) {
	//e := newExpect(t)
	fmt.Println(Mydatabase.QueryUserByName("dxw"))
}

// 创建一个用户
func TestInsertUser(t *testing.T) {
	// e := newExpect(t)
	user := Mydatabase.Userinfo{
		Id: 0, Name: "dxw", FollowCount: 10,
		FollowerCount: 20, Avator: "哈实习",
		BackgroundImage: "赫斯", Signature: "早上好",
		TotalFavorited: 87,
		WorkCount:      23,
		FavoriteCount:  67}
	res := Mydatabase.InsertUser(&user)

	fmt.Println(res)
}

func TestUpdateUser(t *testing.T) {
	// e := newExpect(t)
	user := Mydatabase.Userinfo{
		Id: 90, Name: "大哥2", FollowCount: 10,
		FollowerCount: 20, Avator: "哈实习",
		BackgroundImage: "赫斯", Signature: "早上好",
		TotalFavorited: 87,
		WorkCount:      23,
		FavoriteCount:  67}
	res := Mydatabase.UpdateUser(&user)

	fmt.Println(res)
}
