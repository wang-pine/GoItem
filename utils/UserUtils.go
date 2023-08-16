package utils

import (
	"Mydatabase"
	"strconv"
)

func GetUserToken(Id int64) string {
	password := Mydatabase.QueryUserPWD(Id)
	return strconv.Itoa(int(Id)) + password
}
