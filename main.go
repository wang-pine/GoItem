package main

import (
	"github.com/gin-gonic/gin"
	"service"
	"tools"
)

func main() {
	go service.RunMessageServer()

	r := gin.Default()
	tools.InitRouter(r)

	err := r.Run(":8888")
	if err != nil {
		return
	}
}
