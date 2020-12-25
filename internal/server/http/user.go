package http

import (
	"github.com/gin-gonic/gin"
	"shippo-server/utils/box"
)

func initUserRouter(Router *gin.RouterGroup) {
	user := Router.Group("user")
	{
		user.GET("login", box.Handler(userLogin))
	}
}

func userLogin(c *box.Context) {
	c.JSON(gin.H{
		"message": "你好，世界",
	}, nil)
}
