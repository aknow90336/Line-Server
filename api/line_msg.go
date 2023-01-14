package api

import (
	"MessagingServer/domain/req"
	"MessagingServer/env"
	messageImpl "MessagingServer/repo/impl"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/spf13/viper"
)

func GetUserMsg(c *gin.Context) {
	ctx := c.Request.Context()
	userIdStr := c.Query("user_id")
	repo := messageImpl.NewMessageRepo(ctx, env.InitMongo(ctx))
	res, err := repo.GetMessageByUserId(userIdStr)
	if err != nil {
		log.Println(err)
		_ = c.Error(err)
		return
	}
	c.JSON(200, res)
}

func PushMsg(c *gin.Context) {
	req := &req.Message{}
	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.Error(err)
		return
	}
	bot, err := linebot.New(
		viper.GetString("line.secret"),
		viper.GetString("line.token"),
	)
	if err != nil {
		log.Println(err)
	}
	if _, err := bot.PushMessage(req.UserId, linebot.NewTextMessage(req.Message)).Do(); err != nil {
		log.Print(err)
	}

	c.JSON(200, "success")
}
