package test

import (
	"Mydatabase"
	"fmt"
	"testing"
)

func TestInitUsersDatabase(t *testing.T) {
	err := Mydatabase.InitUsersDatabase()
	if err != nil {
		fmt.Println("init DB failed,err \n", err)
		t.Errorf("发生错误")
	}
}
func TestMakeNewUserTable(t *testing.T) {
	err := Mydatabase.MakeNewUserTable(2)
	if err != nil {
		fmt.Println("init DB failed,err \n", err)
		t.Errorf("发生错误")
	}
}
func TestInsertVideoIdToUserTable(t *testing.T) {
	Mydatabase.InsertVideoIdToUserTable(5,2)
	Mydatabase.InsertVideoIdToUserTable(6,2)
	Mydatabase.InsertVideoIdToUserTable(7,2)
}
func TestGetUserVideoSList(t *testing.T){
	ret,len := Mydatabase.GetUserVideosList(2)
	fmt.Println("ret的长度是",len)
	var i int
	for i = 0; i < len ;i ++{
		fmt.Println(ret[i])
	}
}