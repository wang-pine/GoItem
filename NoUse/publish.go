package nouse

// Publish check token then save upload file to public directory
//func Publish(c *gin.Context) {
//	token := c.PostForm("token")
//
//	if _, exist := usersLoginInfo[token]; !exist {
//		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
//		return
//	}
//
//	data, err := c.FormFile("data")
//	if err != nil {
//		c.JSON(http.StatusOK, common.Response{
//			StatusCode: 1,
//			StatusMsg:  err.Error(),
//		})
//		return
//	}
//
//	filename := filepath.Base(data.Filename)
//	user := usersLoginInfo[token]
//	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
//	saveFile := filepath.Join("./public/", finalName)
//	if err := c.SaveUploadedFile(data, saveFile); err != nil {
//		c.JSON(http.StatusOK, common.Response{
//			StatusCode: 1,
//			StatusMsg:  err.Error(),
//		})
//		return
//	}
//
//	c.JSON(http.StatusOK, common.Response{
//		StatusCode: 0,
//		StatusMsg:  finalName + " uploaded successfully",
//	})
//}
