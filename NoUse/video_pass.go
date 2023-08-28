package nouse

//"github.com/your/package/douyinpb"

/*
func main() {

douyinFeedRequest := &douyinpb.DouyinFeedRequest{
		LatestTime: 0,
		Token:      "your_token_here",
	}


	responseData := SendDouyinFeedRequest(douyinFeedRequest)

	douyinFeedResponse := &douyinpb.DouyinFeedResponse{}
	err := douyinFeedResponse.Unmarshal(responseData)
	if err != nil {
		fmt.Println("Error parsing response:", err)
		return
	}

	// 模拟生成 video_list 数据
	videoList := GenerateVideoList()

	// 对视频列表按照投稿时间进行倒序排序
	sort.Slice(videoList, func(i, j int) bool {
		return videoList[i].GetLatestTime() > videoList[j].GetLatestTime()
	})

	// 获取前 30 个视频（单次最多30个）
	limitedVideoList := videoList
	if len(videoList) > 30 {
		limitedVideoList = videoList[:30]
	}

	response := &douyinpb.DouyinFeedResponse{
		StatusCode: douyinpb.DouyinFeedResponse_Success,
		StatusMsg:  "Success",
		VideoList:  limitedVideoList,
		NextTime:   limitedVideoList[len(limitedVideoList)-1].LatestTime,
	}

	// 打印视频信息
	for _, video := range response.GetVideoList() {
		fmt.Println("Video ID:", video.GetId())
		fmt.Println("Author:", video.GetAuthor().GetName())
		fmt.Println("Play URL:", video.GetPlayUrl())
		fmt.Println("Cover URL:", video.GetCoverUrl())
		fmt.Println("Favorite Count:", video.GetFavoriteCount())
		fmt.Println("Comment Count:", video.GetCommentCount())
		fmt.Println("Title:", video.GetTitle())
		// ...
	}
}
/*
func SendDouyinFeedRequest(request *douyinpb.DouyinFeedRequest) []byte {

	var responseData []byte
	return responseData
}

func GenerateVideoList() []*douyinpb.Video {
	// 模拟生成 video_list 数据
	videoList := []*douyinpb.Video{
		{
			Id:            123456,
			Author:        &douyinpb.User{Id: 789, Name: "User1", IsFollow: false},
			PlayUrl:       "http://example.com/video/123456",
			CoverUrl:      "http://example.com/cover/123456",
			FavoriteCount: 100,
			CommentCount:  50,
			IsFavorite:    false,
			Title:         "Video 1",
			LatestTime:    1679000000,
		},
		// 添加更多视频...
	}

	return videoList
}
*/
