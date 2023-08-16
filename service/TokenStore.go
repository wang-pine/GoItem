package service

//做一个哈希表，用于对Token进行维护
//注意，这里可以使用redis进行替代
//哈希表只是权宜之策
var tokenStore = make(map[string]int64)
var idStore = make(map[int64]string)

//"token":userID
//userId:"token"
//维护两个哈希表实现键值对的快速查询
func PushToken(token string, userId int64) bool {
	if tokenStore[token] == 0 {
		tokenStore[token] = userId
		idStore[userId] = token
	} else {
		return false
	}
	return true
}

func SearchToken(token string) (ok bool, userId int64) {
	if tokenStore[token] == 0 {
		return false, 0
	} else {
		return true, tokenStore[token]
	}
}
func SearchTokenById(Id int64) (ok bool, token string) {
	if idStore[Id] == "" {
		return false, ""
	} else {
		return true, idStore[Id]
	}
}
