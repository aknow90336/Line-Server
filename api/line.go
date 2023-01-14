package api

import (
	"github.com/gin-gonic/gin"
)

func CallBack(c *gin.Context) {

	c.JSON(200, nil)
}
