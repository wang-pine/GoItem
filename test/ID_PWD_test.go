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
	err := Mydatabase.InsertNewUser("wsasnan")
	if err != nil {
		fmt.Println("init DB failed,err \n", err)
		t.Errorf("发生错误")
	}
}