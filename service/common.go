package service

type Response struct {
	StatusCode int32
	StatusMsg  string
}

type Video struct {
	Id            int64
	Author        User
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
	IsFavorite    bool
}

type Comment struct {
	Id         int64
	User       User
	Content    string
	CreateDate string
}

type User struct {
	Id            int64
	Name          string
	FollowCount   int64
	FollowerCount int64
	IsFollow      bool
}

type Message struct {
	Id         int64
	Content    string
	CreateTime string
}
type MessageSendEvent struct {
	UserId     int64
	ToUserId   int64
	MsgContent string
}

type MessagePushEvent struct {
	FromUserId int64
	MsgContent string
}
