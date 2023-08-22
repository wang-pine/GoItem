package test

import (
	"Mydatabase"
	"testing"
)

func TestMakeCommentTable(t *testing.T) {
	Mydatabase.MakeCommentTable(3)
}
func TestInsertComment(t *testing.T) {
	Mydatabase.InsertComment(1, 2, "哈哈哈")
}
func TestDeleteComment(t *testing.T) {
	Mydatabase.DeleteComment(1,4)
}
func TestGetCommentList(t *testing.T) {
	Mydatabase.GetCommentList(1)
}