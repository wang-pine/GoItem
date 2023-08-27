package service

import (
	"Mydatabase/util"
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

// 做一个哈希表，用于对Token进行维护
// 注意，这里可以使用redis进行替代
// 哈希表只是权宜之策
var tokenStore = make(map[string]int64)
var idStore = make(map[int64]string)

// "token":userID
// userId:"token"
// 维护两个哈希表实现键值对的快速查询
func PushToken(token string, userId int64) bool {
	userIdstr := strconv.FormatInt(userId, 10)
	// if tokenStore[token] == 0 {
	// 	tokenStore[token] = userId
	// 	idStore[userId] = token
	// } else {
	// 	return false
	// }
	// return true
	tag, _ := util.Get("token")
	if tag == false {
		util.Set(token, userIdstr)
		util.Set(userIdstr, token)
	} else {
		return false
	}
	return true

}

func SearchToken(token string) (ok bool, userId int64) {
	tag, res := util.Get(token)
	nt64, err := strconv.ParseInt(res, 10, 64)
	if tag == false || err != nil {
		return false, 0
	} else {
		return true, nt64
	}
}
func SearchTokenById(Id int64) (ok bool, token string) {
	tag, res := util.Get(strconv.FormatInt(Id, 10))
	if tag == false {
		return false, ""
	} else {
		return true, res
	}
	//if idStore[Id] == "" {
	//	return false, ""
	//} else {
	//	return true, idStore[Id]
	//}
}

// MD5加密
func StringToMD5(PWD string) string {
	w := md5.New()
	w.Write([]byte(PWD))
	return hex.EncodeToString(w.Sum(nil))
}
func CreateUserToken(Id int64, password string) string {
	token := strconv.FormatInt(Id, 10) + StringToMD5(password)

	return token
}
