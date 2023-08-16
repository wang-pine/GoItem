package Mydatabase

//用来对基本信息进行维护
import (
	"common"
	"fmt"
)

// 用户信息结构体
/*
type Userinfo struct {
	Id              int64  `gorm:"type:varchar(20); not null" json:"id" binding:"required"`
	Name            string `gorm:"type:varchar(20); not null" json:"name" binding:"required"`
	FollowCount     int64  `gorm:"type:int(64); not null" json:"follow_count" binding:"required"`
	FollowerCount   int64  `gorm:"type:int(64); not null" json:"follower_count" binding:"required"`
	Avator          string `gorm:"type:varchar(256); not null" json:"avator" binding:"required"`
	BackgroundImage string `gorm:"type:varchar(256); not null" json:"background_image" binding:"required"`
	Signature       string `gorm:"type:varchar(64); not null" json:"signature" binding:"required"`
	TotalFavorited  int64  `gorm:"type:int(20); not null" json:"total_favorited" binding:"required"`
	WorkCount       int64  `gorm:"type:int(20); not null" json:"work_count" binding:"required"`
	FavoriteCount   int64  `gorm:"type:int(20); not null" json:"favorite_count" binding:"required"`
}
*/
// 通过id查询userinfo信息
func QueryUserById(id int64) []common.Userinfo {
	db, err := GetDB()
	var users []common.Userinfo
	db.Where("id = ?", id).Find(&users)
	if err != nil {
		fmt.Println("连接失败！！")
	}
	return users
}

// 通过name查询userinfo信息
func QueryUserByName(name string) []common.Userinfo {
	db, err := GetDB()
	var users []common.Userinfo
	db.Where("name = ?", name).Find(&users)
	if err != nil {
		fmt.Println("连接失败！！")
	}
	return users
}

// 插入一个用户,1成功，-失败
func InsertUser(user *common.Userinfo) int64 {
	db, err := GetDB()
	if err != nil {
		fmt.Println("连接失败！！")
		
	}
	result := db.Create(&user)
	if result.Error != nil {
		return -1
	}
	return 1
}

// 修改一个用户,1成功，-失败
func UpdateUser(user *common.Userinfo) int64 {
	db, err := GetDB()
	if err != nil {
		fmt.Println("连接失败！！")
		return -1
	}
	id := user.Id
	var users []common.Userinfo
	db.Where("id = ?", id).Find(&users)
	if len(users) == 0 {
		return -1
	}
	db.Where("id = ?", id).Save(&user)
	return 1
}
