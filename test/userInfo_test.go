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
func TestMakeUserTable(t *testing.T) {
	err := Mydatabase.MakeNewUserTable(2)
	if err != nil {
		fmt.Println("init DB failed,err \n", err)
		t.Errorf("发生错误")
	}
}
