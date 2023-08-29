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

	video.Author.Avatar = videoInfo.AuthorAvator
	video.Author.BackgroundImage = videoInfo.AuthorBackgroundImage
	video.Author.Signature = videoInfo.AuthorSignature
	video.Author.TotalFavorited = videoInfo.AuthorTotalFavorited
	video.Author.WorkCount = videoInfo.AuthorWorkCount
	video.Author.FavoriteCount = videoInfo.AuthorFavoriteCount

	video.PlayUrl = videoInfo.VideoPlayUrl
	video.CoverUrl = videoInfo.VideoCoverUrl
	video.FavoriteCount = videoInfo.VideoFavoriteCount
	video.CommentCount = videoInfo.VideoCommentCount
	video.IsFavorite = Mydatabase.IsFavorite(userId, video.Id)
	video.Title = videoInfo.VideoTitle
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
	videoInfo.AuthorAvator = video.Author.Avatar
	videoInfo.AuthorBackgroundImage = video.Author.BackgroundImage
	videoInfo.AuthorSignature = video.Author.Signature
	videoInfo.AuthorTotalFavorited = video.Author.TotalFavorited
	videoInfo.AuthorWorkCount = video.Author.WorkCount
	videoInfo.AuthorFavoriteCount = video.Author.FavoriteCount
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
	user.Avatar = userInfo.Avator
	user.BackgroundImage = userInfo.BackgroundImage
	user.Signature = userInfo.Signature
	user.TotalFavorited = userInfo.TotalFavorited
	user.WorkCount = userInfo.WorkCount
	user.FavoriteCount = userInfo.FavoriteCount
}

// user是前端向后端传递的信息
// userinfo是向数据库传递的信息
// 这个函数是向数据库传递信息用的
func ConvertUserToUserIfo(user *common.User, userInfo *common.Userinfo) {
	userInfo.Id = user.Id
	userInfo.Name = user.Name
	userInfo.FollowCount = user.FollowCount
	userInfo.FollowerCount = user.FollowerCount
	userInfo.Avator = user.Avatar
	userInfo.BackgroundImage = user.BackgroundImage
	userInfo.Signature = user.Signature
	userInfo.TotalFavorited = user.TotalFavorited
	userInfo.WorkCount = user.WorkCount
	userInfo.FavoriteCount = user.FavoriteCount
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
