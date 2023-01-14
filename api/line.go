package api

import (
	"MessagingServer/env"
	messageImpl "MessagingServer/repo/impl"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/spf13/viper"
)

func CallBack(c *gin.Context) {

	ctx := c.Request.Context()

	bot, err := linebot.New(
		viper.GetString("line.secret"),
		viper.GetString("line.token"),
	)

	if err != nil {
		log.Fatal(err)
	}
	events, err := bot.ParseRequest(c.Request)
	if err != nil {
		log.Fatal(err)
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				repo := messageImpl.NewMessageRepo(ctx, env.InitMongo(ctx))
				err := repo.AddMessage(event.Source.UserID, message.Text)
				if err != nil {
					log.Fatal(err)
				}

			}
		}
	}

	c.JSON(200, nil)
}
