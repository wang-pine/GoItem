package Mydatabase
/*
********************
存储视频完整信息
********************
*/
//这里用来对视频相关的数据库进行维护
import (
	"common"
	"fmt"
)

// 通过视频id查询Videoinfo信息
func QueryVideoById(id int64) common.Videoinfo {
	db, err := GetDB()
	var videos common.Videoinfo
	db.Where("video_id = ?", id).Find(&videos)
	if err != nil {
		fmt.Println("连接失败！！")
	}
	return videos
}

// 通过作者姓名查询Videoinfo信息
func QueryVideoByAuthorName(AuthorName string) []common.Videoinfo {
	db, err := GetDB()
	var videos []common.Videoinfo
	db.Where("author_name = ?", AuthorName).Find(&videos)
	if err != nil {
		fmt.Println("连接失败！！")
	}
	return videos
}

// 通过视频标题查询Videoinfo信息
func QueryVideoByVideoTitle(VideoTitle string) []common.Videoinfo {
	db, err := GetDB()
	var videos []common.Videoinfo
	db.Where("video_title = ?", VideoTitle).Find(&videos)
	if err != nil {
		fmt.Println("连接失败！！")
	}
	return videos
}

// 通过用户id查询Videoinfo信息
func QueryVideoByAuthorId(authorId int64) []common.Videoinfo {
	db, err := GetDB()
	var videos []common.Videoinfo
	db.Where("author_id = ?", authorId).Find(&videos)
	if err != nil {
		fmt.Println("连接失败！！")
	}
	return videos
}

// 通过用户id查询该用户所有视频id
func QueryVideoIdByAuthorId(authorId int64) []int64 {
	db, err := GetDB()
	var videos []common.Videoinfo
	db.Where("author_id = ?", authorId).Find(&videos)
	if err != nil {
		fmt.Println("连接失败！！")
	}
	vedisIds := make([]int64, len(videos))
	for i := 0; i < len(videos); i++ {
		vedisIds[i] = videos[i].VideoId
	}
	return vedisIds
}

// 通过用户姓名查询该用户所有视频id
func QueryVideoIdByAuthorName(authorName string) []int64 {
	db, err := GetDB()
	var videos []common.Videoinfo
	db.Where("author_name = ?", authorName).Find(&videos)
	if err != nil {
		fmt.Println("连接失败！！")
	}
	vedisIds := make([]int64, len(videos))
	for i := 0; i < len(videos); i++ {
		vedisIds[i] = videos[i].VideoId
	}
	return vedisIds
}

// 通过video_id删除vedio
func DeleteByVideoId(videoId int64) bool {
	db, err := GetDB()
	var videos []common.Videoinfo
	db.Where("video_id = ?", videoId).Find(&videos)
	if err != nil {
		fmt.Println("连接失败！！")
	}
	if len(videos) == 0 {
		return false
	}
	db.Where("author_id = ?", videoId).Delete(&videos)
	return true
}

// 通过用户id删除vedio
func DeleteByAuthorId(authorId int64) bool {
	db, err := GetDB()
	var videos []common.Videoinfo
	db.Where("author_id = ?", authorId).Find(&videos)
	if err != nil {
		fmt.Println("连接失败！！")
	}
	if len(videos) == 0 {
		return false
	}

	db.Where("author_id = ?", authorId).Delete(&videos)
	return true
}

// 增加视频信息
func InsertVideoInfo(video *common.Videoinfo) bool {
	db, err := GetDB()
	if err != nil {
		fmt.Println("连接失败！！")
	}
	result := db.Create(&video)
	if result.Error != nil {
		return false
	}
	return true
}

// 修改视频信息
func UpdateVideoInfo(video *common.Videoinfo) bool {
	db, err := GetDB()
	var videos []common.Videoinfo
	videoId := video.VideoId
	db.Where("video_id = ?", videoId).Find(&videos)
	if err != nil {
		fmt.Println("连接失败！！")
	}
	fmt.Println(videos)
	if len(videos) == 0 {
		return false
	}
	db.Where("video_id = ?", videoId).Save(&video)
	return true
}

// 获取最后一个视频
func GetLastVideo() common.Videoinfo {
	db, err := GetDB()
	if err != nil {
		fmt.Println("链接数据库失败")
	}
	var video common.Videoinfo
	db.Last(&video)
	return video
}
