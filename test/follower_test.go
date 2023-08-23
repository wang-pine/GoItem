package test

import (
	"Mydatabase"
	"fmt"
	"testing"
)
func TestInitFollowersFatabbase(t *testing.T) {
	Mydatabase.InitFollowersDatabase()
}
func TestMakeNewFollowerTable(t *testing.T) {
	Mydatabase.MakeNewFollowerTable(2)
}
func TestInsertFollowerIdToUserTable(t *testing.T) {
	Mydatabase.InsertFollowerIdToUserTable(17,2)
}
func TestGetUserFollowersList(t *testing.T) {
	list,len:=Mydatabase.GetUserFollowersList(2)
	var i int
	for i=0;i<len;i++{
		fmt.Println(list[i])
	}
}
func TestIsFollow(t *testing.T) {
	fmt.Println(Mydatabase.IsFollow(2,11))
	fmt.Println(Mydatabase.IsFollow(2,12))
}
func TestDeleteFollower(t *testing.T) {
	Mydatabase.DeleteFollower(3,1)
}