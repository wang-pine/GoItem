package test

import (
	"Mydatabase"
	"fmt"
	"testing"
)

func TestInitInfoDatabase(t *testing.T) {
	//e := newExpect(t)
	err := Mydatabase.InitInfoDatabase()
	if err != nil {
		fmt.Println("init DB failed,err \n", err)
	}
	ret := Mydatabase.InsertPWD("'wangs'")
	fmt.Println("此时的id是", ret)
	if ret != 2 {
		t.Errorf("id理应是2但是现在是:%v", ret)
	}
}
