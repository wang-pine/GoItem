package test

import (
	"Mydatabase"
	"fmt"
	"testing"
)

func TestInitVideosDatabase(t *testing.T) {
	err := Mydatabase.InitVideosDatabase()
	if err != nil {
		fmt.Println("init DB failed,err \n", err)
		t.Errorf("发生错误")
	}
}
func TestMakeNewVideoTable(t *testing.T) {
	err0 := Mydatabase.MakeNewVideoTable(1)
	if err0 != nil {
		fmt.Println("make table failed,err \n", err0)
		t.Errorf("发生错误")
	}
	err1 := Mydatabase.MakeNewVideoTable(2)
	if err1 != nil {
		fmt.Println("make table failed,err \n", err1)
		t.Errorf("发生错误")
	}
	err2 := Mydatabase.MakeNewVideoTable(3)
	if err2 != nil {
		fmt.Println("make table failed,err \n", err2)
		t.Errorf("发生错误")
	}
}
func TestInsertUserIdToVideoTable(t *testing.T) {
	Mydatabase.InsertUserIdToVideoTable(1, 3)
	Mydatabase.InsertUserIdToVideoTable(2, 3)
	Mydatabase.InsertUserIdToVideoTable(3, 4)
	Mydatabase.InsertUserIdToVideoTable(1, 4)
}
func TestGetFavoriteUsersList(t *testing.T) {
	ret, len := Mydatabase.GetFavoriteUsersList(1)
	var i int
	for i = 0; i < len; i++ {
		fmt.Println(ret[i], " ")
	}
}
func TestIsFavorite(t *testing.T){
	fmt.Println(Mydatabase.IsFavorite(4,1))
	fmt.Println(Mydatabase.IsFavorite(5,1))
}