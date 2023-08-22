package test

import (
	"Mydatabase"
	"fmt"
	"testing"
)

func TestInitFavoriteDatabase(t *testing.T) {
	err := Mydatabase.InitFavoriteDatabase()
	if err != nil {
		fmt.Println("init DB failed,err \n", err)
		t.Errorf("发生错误")
	}
}

func TestFavoriteMake(t *testing.T) {
	err := Mydatabase.MakeNewFavoriteTable(5)
	if err != nil {
		fmt.Println("init DB failed,err \n", err)
		t.Errorf("发生错误")
	}
}

func TestFavoriteInsert(t *testing.T) {
	Mydatabase.InsertUserIdToFavoriteTable(5, 3)

}

func TestFavoriteDelete(t *testing.T) {
	Mydatabase.DeleteUserIdToFavoriteTable(5, 2)

}

func TestFavoriteQuery(t *testing.T) {
	fmt.Println(Mydatabase.GetFavoriteVideoList(5))

}
