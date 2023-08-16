package NoUse
/*
	func Register(c *gin.Context) {
		username := c.Query("username")
		password := c.Query("password")

		token := username + password

		if _, exist := usersLoginInfo[token]; exist {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: common.Response{StatusCode: 1, StatusMsg: "User already exist"},
			})
		} else {
			atomic.AddInt64(&userIdSequence, 1)
			newUser := common.User{
				Id:   userIdSequence,
				Name: username,
			}
			usersLoginInfo[token] = newUser
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: common.Response{StatusCode: 0},
				UserId:   userIdSequence,
				Token:    username + password,
			})
		}
	}

	func Login(c *gin.Context) {
		username := c.Query("username")
		password := c.Query("password")

		token := username + password

		if user, exist := usersLoginInfo[token]; exist {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: common.Response{StatusCode: 0},
				UserId:   user.Id,
				Token:    token,
			})
		} else {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
			})
		}
	}
*/