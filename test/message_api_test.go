package test

import (
	"Mydatabase"
	"fmt"
	"testing"
)

func TestInitMessageDatabase(t *testing.T) {
	Mydatabase.InitMessageDatabase()
}
func TestMakeMessageTable(t *testing.T) {
	Mydatabase.MakeNewMessageTable(4)
}
func TestInsertMessage(t *testing.T) {
	Mydatabase.InsertMessage(1, 2, "大家好哇")
}
func TestMessageList(t *testing.T) {
	messageList := Mydatabase.GetMessageList(2, 0)
	length := len(messageList)
	var i int
	for i = 0; i < length; i++ {
		fmt.Println(messageList[i])
	}
}
