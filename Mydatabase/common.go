package Mydatabase

//这个是用在与数据库进行交互的结构体
//Video从数据库查询
//（id是主键）
type VideoData struct {
	Id            int64
	Author        UserData
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
	UploadTime    int64
	//老实说，用int64存时间戳是不安全的，再想办法
}

//用户信息从数据库查询
//(id是主键)
type UserData struct {
	Id               int64
	Name             string
	FollowCount      int64
	FollowerCount    int64
	avator           string
	backgroundImage  string
	signature        string
	tootal_favorited int64
	work_count       int64
	favorite_count   int64
}

//用户信息
//用户的账号和密码
//id和pwd从数据库查询
type IDPWD struct {
	Id  int64
	PWD string
}

//用户累计上传视频的结构体
//这个键值对在用户分表中存储
type UserVideo struct {
	userId  int64
	videoId int64
}

//视频累计点赞的结构体
//这个键值对在视频分表中存储
type VideoFavorited struct {
	videoId int64
	userId  int64
}
