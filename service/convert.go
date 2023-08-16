// 这个文件用于数据库信息和上传到前端信息的结构体进行对接
package service

import (
	"Mydatabase"
	"common"
	//"controller"
)

// videoinfo是数据库信息
// video是传向前端进行渲染的信息
// userId是当前登录的用户的id，用来检查是否有关注这个视频的博主
func ConvertVideoInfoToVideo(videoInfo *common.Videoinfo, video *common.Video, userId int64) {
	video.Id = videoInfo.VideoId
	video.Author.Id = videoInfo.AuthorId
	video.Author.Name = videoInfo.AuthorName
	video.Author.FollowCount = videoInfo.AuthorFollowCount
	video.Author.FollowerCount = videoInfo.AuthorFollowerCount
	video.Author.IsFollow = Mydatabase.IsFollow(video.Author.Id, userId)
	video.PlayUrl = videoInfo.VideoPlayUrl
	video.CoverUrl = videoInfo.VideoCoverUrl
	video.FavoriteCount = videoInfo.VideoFavoriteCount
	video.CommentCount = videoInfo.VideoCommentCount
	video.IsFavorite = Mydatabase.IsFavorite(userId, video.Id)
}

// 将video信息存入数据库中
// video是从前端接受的信息
// videoinfo是发往数据库的信息
func ConvertVideoToVideoInfo(video *common.Video, videoInfo *common.Videoinfo) {
	videoInfo.VideoId = video.Id
	videoInfo.AuthorId = video.Author.Id
	videoInfo.AuthorName = video.Author.Name
	videoInfo.AuthorFollowCount = video.Author.FollowCount
	videoInfo.AuthorFollowerCount = video.Author.FollowerCount
	videoInfo.AuthorAvator = ""
	videoInfo.AuthorBackgroundImage = ""
	videoInfo.AuthorSignature = ""
	videoInfo.AuthorTotalFavorited = 0
	videoInfo.AuthorWorkCount = 0
	videoInfo.AuthorFavoriteCount = 0
	videoInfo.VideoPlayUrl = video.PlayUrl
	videoInfo.VideoCoverUrl = video.CoverUrl
	videoInfo.VideoFavoriteCount = video.FavoriteCount
	videoInfo.VideoCommentCount = video.CommentCount
}

// userinfo是数据库结构体
// user是向前端传递的信息结构体
// 这个函数是向前端传递信息用的
func ConvertUserInfoToUser(userInfo *common.Userinfo, user *common.User, userId int64) {
	user.Id = userInfo.Id
	user.Name = userInfo.Name
	user.FollowCount = userInfo.FollowCount
	user.FollowerCount = userInfo.FollowerCount
	user.IsFollow = Mydatabase.IsFollow(user.Id, userId)
}

// user是前端向后端传递的信息
// userinfo是向数据库传递的信息
// 这个函数是向数据库传递信息用的
func ConvertUserToUserIfo(user *common.User, userInfo *common.Userinfo) {
	userInfo.Id = user.Id
	userInfo.Name = user.Name
	userInfo.FollowCount = user.FollowCount
	userInfo.FollowerCount = user.FollowerCount
	userInfo.Avator = ""
	userInfo.BackgroundImage = ""
	userInfo.Signature = ""
	userInfo.TotalFavorited = 0
	userInfo.WorkCount = 0
	userInfo.FavoriteCount = 0
}

// user是前端向后端传递的信息
// userinfo是向数据库传递的信息
// 这个函数是向数据库传递信息用的
func ConvertUserVideoToVideoIfo(userInfo *common.Userinfo, video *common.Video, videoInfo *common.Videoinfo) {
	videoInfo.VideoId = video.Id
	videoInfo.AuthorId = video.Author.Id
	videoInfo.AuthorName = video.Author.Name
	videoInfo.AuthorFollowCount = video.Author.FollowCount
	videoInfo.AuthorFollowerCount = video.Author.FollowerCount
	videoInfo.AuthorAvator = userInfo.Avator
	videoInfo.AuthorBackgroundImage = userInfo.BackgroundImage
	videoInfo.AuthorSignature = userInfo.Signature
	videoInfo.AuthorTotalFavorited = userInfo.TotalFavorited
	videoInfo.AuthorWorkCount = userInfo.WorkCount
	videoInfo.AuthorFavoriteCount = userInfo.FavoriteCount
	videoInfo.VideoPlayUrl = video.PlayUrl
	videoInfo.VideoCoverUrl = video.CoverUrl
	videoInfo.VideoFavoriteCount = video.FavoriteCount
	videoInfo.VideoCommentCount = video.CommentCount
}
