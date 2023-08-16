package test

import (
	"Mydatabase"
	"fmt"
	"testing"
)

func TestInitPWDDatabase(t *testing.T) {
	err := Mydatabase.InitPWDDatabase()
	if err != nil {
		fmt.Println("init DB failed,err \n", err)
		t.Errorf("发生错误")
	}
}
func TestInsertNewUser(t *testing.T) {
	err,_ := Mydatabase.InsertNewUser("wodiaonimade")
	if err != nil {
		fmt.Println("init DB failed,err \n", err)
		t.Errorf("发生错误")
	}
}
func TestQueryUserPWD(t *testing.T) {
	pwd:=Mydatabase.QueryUserPWD(2)
	fmt.Println(pwd)
}
func TestJudgePWD(t *testing.T) {
	test :="wsasnan"
	ok := Mydatabase.JudgePWD(2,test)
	fmt.Println(ok)
}