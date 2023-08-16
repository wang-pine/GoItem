package test

import (
	"Mydatabase"
	"fmt"
	"math/rand"
	"testing"
)

func TestQueryVideoById(t *testing.T) {
	fmt.Println(Mydatabase.QueryVideoById(1))
}

func TestQueryVideoByAuthorId(t *testing.T) {
	fmt.Println(Mydatabase.QueryVideoByAuthorId(2))
}

func TestInsertVideoInfo(t *testing.T) {
	video := Mydatabase.Videoinfo{
		VideoId: 2, AuthorId: 2, AuthorName: "测试数据",
		AuthorFollowCount: 21, AuthorFollowerCount: 23,
		AuthorAvator: "eferg", AuthorBackgroundImage: "深夜下哦美好看",
		AuthorSignature: "qg", AuthorTotalFavorited: 21, AuthorWorkCount: 765,
		AuthorFavoriteCount: 43, VideoPlayUrl: "更好的武功和",
		VideoCoverUrl: "ge", VideoFavoriteCount: 231,
		VideoCommentCount: 233, VideoTitle: "好看的时间噢批",
		VideoTime: "geq",
	}
	res := Mydatabase.InsertVideoInfo(&video)
	fmt.Println(res)
}

func TestQueryVideoByAuthorName(t *testing.T) {
	fmt.Println(Mydatabase.QueryVideoByAuthorName("测试数据"))
}

func TestQueryVideoByVideoTitle(t *testing.T) {
	fmt.Println(Mydatabase.QueryVideoByVideoTitle("哈哈哈哈"))
}

func TestQueryVideoIdByAuthorId(t *testing.T) {
	fmt.Println(Mydatabase.QueryVideoIdByAuthorId(1))
}

func TestQueryVideoIdByAuthorName(t *testing.T) {
	fmt.Println(Mydatabase.QueryVideoIdByAuthorName("测试数据"))
}

func TestDeleteByAuthorId(t *testing.T) {
	//注意，不要测试delete，没有授权
	fmt.Println(Mydatabase.DeleteByAuthorId(1))
}

func TestUpdateVideoInfo(t *testing.T) {
	video := Mydatabase.Videoinfo{
		VideoId: 5, AuthorId: 2, AuthorName: "垃圾很多很多ddddd",

		AuthorFollowCount: 21, AuthorFollowerCount: 23,
		AuthorAvator: "的结局", AuthorBackgroundImage: "d334",
		AuthorSignature: "二分", AuthorTotalFavorited: 21, AuthorWorkCount: 765,
		AuthorFavoriteCount: 43, VideoPlayUrl: "ddfa",
		VideoCoverUrl: "而非我", VideoFavoriteCount: 231,
		VideoCommentCount: 233, VideoTitle: "dfwfw的我",
		VideoTime: "微服务",
	}
	fmt.Println(Mydatabase.UpdateVideoInfo(&video))
}
func TestDeleteByVideoId(t *testing.T) {
	fmt.Println(Mydatabase.DeleteByVideoId(2))
}

func TestRand(t *testing.T) {
	fmt.Println(rand.Int())
}
