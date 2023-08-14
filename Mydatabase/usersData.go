package Mydatabase

//用来对基本信息进行维护

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// 用户信息结构体
type Userinfo struct {
	Id              int64  `gorm:"type:varchar(20); not null" json:"id" binding:"required"`
	Name            string `gorm:"type:varchar(20); not null" json:"name" binding:"required"`
	FollowCount     int64  `gorm:"type:int(64); not null" json:"followcount" binding:"required"`
	FollowerCount   int64  `gorm:"type:int(64); not null" json:"followercount" binding:"required"`
	Avator          string `gorm:"type:varchar(256); not null" json:"avator" binding:"required"`
	BackgroundImage string `gorm:"type:varchar(256); not null" json:"backgroundimage" binding:"required"`
	Signature       string `gorm:"type:varchar(64); not null" json:"signature" binding:"required"`
	TotalFavorited  int64  `gorm:"type:int(20); not null" json:"totalfavorited" binding:"required"`
	WorkCount       int64  `gorm:"type:int(20); not null" json:"workcount" binding:"required"`
	FavoriteCount   int64  `gorm:"type:int(20); not null" json:"favoritecount" binding:"required"`
}

func GetDB() (*gorm.DB, error) {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "douyin:123456@tcp(127.0.0.1:3306)/douyin_info"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	// 连接池
	sqlDB, err := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second) // 10秒钟

	return db, err
}

// 通过id查询userinfo信息
func QueryUserById(id int64) []Userinfo {
	db, err := GetDB()
	var users []Userinfo
	db.Where("id = ?", id).Find(&users)
	if err != nil {
		fmt.Println("连接失败！！")
	}
	return users
}

// 通过name查询userinfo信息
func QueryUserByName(name string) []Userinfo {
	db, err := GetDB()
	var users []Userinfo
	db.Where("name = ?", name).Find(&users)
	if err != nil {
		fmt.Println("连接失败！！")
	}
	return users
}

// 插入一个用户,1成功，-失败
func InsertUser(user *Userinfo) int64 {
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
func UpdateUser(user *Userinfo) int64 {
	db, err := GetDB()
	if err != nil {
		fmt.Println("连接失败！！")
	}
	id := user.Id
	var users []Userinfo
	db.Where("id = ?", id).Find(&users)
	if len(users) == 0 {
		return -1
	}
	db.Save(&user)
	return 1
}

