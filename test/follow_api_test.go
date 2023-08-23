package test

import (
	"Mydatabase"
	"fmt"
	"testing"
)

func TestInitFollowDatabase(t *testing.T) {
	Mydatabase.InitFollowDatabase()
}
func TestMakeNewFollowTable(t *testing.T) {
	Mydatabase.MakeNewFollowTable(3)
}
func TestInsertFollowId(t *testing.T) {
	Mydatabase.InsertFollowIdToUserTable(2, 3)
}
func TestUserFollowList(t *testing.T) {
	list, len := Mydatabase.GetUserFollowList(3)
	var i int
	for i = 0; i < len; i++ {
		fmt.Println(list[i])
	}
}
func TestDeleteFollow(t *testing.T) {
	//Mydatabase.DeleteFollow(3, 1)
	//Mydatabase.DeleteFollow(4, 1)
	Mydatabase.GetUserFollowList(1)
}
