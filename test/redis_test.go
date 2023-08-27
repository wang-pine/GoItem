package test

import (
	"Mydatabase/util"
	"fmt"
	"testing"
)

func TestConnect(t *testing.T) {
	util.Set("Username", "uxxxser")
	tag, res := util.Get("Usedrname")
	if tag == false {
		fmt.Print("没有读取到！！")
	}
	fmt.Println(res)
}
