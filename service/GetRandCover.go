package service

import (
	"config"
	"fmt"
	"math/rand"
	"time"
)

// 本函数用于获取随机的视频封面
func GetRandCover() (avatar string) {
	rand.Seed(time.Now().UnixNano())
	var r int = rand.Intn(5)
	switch r {
	case 1:
		fmt.Println("1")
		avatar = config.GetLocalAddr() + "/static/c_1.jpg"
		break
	case 2:
		fmt.Println("2")
		avatar = config.GetLocalAddr() + "/static/c_2.jpg"
		break
	case 3:
		fmt.Println("3")
		avatar = config.GetLocalAddr() + "/static/c_3.jpg"
		break
	default:
		fmt.Println("0")
		avatar = config.GetLocalAddr() + "/static/c_3.jpg"
		break
	}
	return avatar
}
