package main

import (
	"config"
	"tools"

	"github.com/gin-gonic/gin"
)

func main() {
	//go controller.RunMessageServer()
	config.InitConfig()
	r := gin.Default()
	tools.InitRouter(r)
	err := r.Run(":8080")
	if err != nil {
		return
	}

}
