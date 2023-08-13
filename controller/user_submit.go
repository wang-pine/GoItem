package controller

import (
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Credentials struct {
	Account  int64  `json:"account"`
	Password string `json:"password"`
}

func main() {
	r := gin.Default()

	r.POST("/login", func(c *gin.Context) {
		var creds Credentials

		if err := c.ShouldBindJSON(&creds); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 验证账号和密码
		if !isValidAccount(creds.Account) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account"})
			return
		}

		if !isValidPassword(creds.Password) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
			return
		}

		// 后端向数据库请求比对
		// ...

		/*比对账号密码 需要数据库封装好的函数
		if isValidCredentials(creds.Account, creds.Password) {
			c.JSON(http.StatusOK, gin.H{"message": "登录成功"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "账号或密码错误"})
		}*/
	})

	r.Run(":8080")
}

func isValidAccount(account int64) bool {
	// 账号只能由数字组成
	// 这里使用正则表达式验证，确保只包含数字
	re := regexp.MustCompile("^[0-9]+$")
	return re.MatchString(strconv.FormatInt(account, 10))
}

func isValidPassword(password string) bool {
	// 密码可以由字符和数字混合组成
	// 这里使用正则表达式验证，确保只包含字符和数字
	re := regexp.MustCompile("^[a-zA-Z0-9]+$")
	return re.MatchString(password)
}

/*
`isValidCredentials` 函数用于验证输入的账号和密码是否与数据库中的记录匹配。具体来说，它执行以下操作：

1. 根据输入的账号，从数据库中查找是否有与之对应的用户记录。
2. 如果找到了记录，它会将输入的密码与数据库中的密码进行比较，以验证密码是否正确。

func isValidCredentials(account int64, password string) bool {
	var user User
	result := db.Where("account = ?", account).First(&user)
	if result.Error != nil {
		return false
	}
	if user.Password == password {
		return true
	}

	return false
}
*/
