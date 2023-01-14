package api

import (
	"MessagingServer/env"
	messageImpl "MessagingServer/repo/impl"
	"log"

	"github.com/gin-gonic/gin"
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

	// todo get post body user_id data

	// if _, err := bot.PushMessage("", linebot.NewTextMessage("")).Do(); err != nil {
	// 	log.Print(err)
	// }

	c.JSON(200, nil)
}
