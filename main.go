package main

import (
	"MessagingServer/api"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("dev")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	r := gin.Default()
	r.POST("/callback", api.CallBack)
	r.POST("/message", api.PushMsg)
	r.GET("/message", api.GetUserMsg)
	r.Run(":5000")
}
