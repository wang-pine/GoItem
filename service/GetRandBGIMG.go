package service

import (
	"config"
	"fmt"
	"math/rand"
	"time"
)

// 本函数用于获取用户随机的背景图片
func GetRandBGIMG() (img string) {
	rand.Seed(time.Now().UnixNano())
	var r int = rand.Intn(5)
	switch r {
	case 1:
		fmt.Println("1")
		img = config.GetLocalAddr() + "/static/1.jpg"
		break
	case 2:
		fmt.Println("2")
		img = config.GetLocalAddr() + "/static/2.jpg"
		break
	case 3:
		fmt.Println("3")
		img = config.GetLocalAddr() + "/static/3.jpg"
		break
	default:
		fmt.Println("0")
		img = config.GetLocalAddr() + "/static/3.jpg"
		break
	}
	return img
}
