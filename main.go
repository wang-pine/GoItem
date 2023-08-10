package main

import (
	"service"
	"tools"

	"github.com/gin-gonic/gin"
)

func main() {
	go service.RunMessageServer()

	r := gin.Default()

	tools.InitRouter(r)

	r.Run(":8888")
}
