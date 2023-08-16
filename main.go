package main

import (
	"controller"
	"tools"

	"github.com/gin-gonic/gin"
)

func main() {
	go controller.RunMessageServer()

	r := gin.Default()

	tools.InitRouter(r)

	r.Run(":8888")
}
