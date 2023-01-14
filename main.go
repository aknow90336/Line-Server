package main

import (
	"MessagingServer/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/callback", api.CallBack)

	r.Run(":5000")
}
