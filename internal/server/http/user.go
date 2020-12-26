package http

import (
	"fmt"
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
	var param = new(struct {
		Phone string `json:"phone"`
		Code  string `json:"code"`
	})
	c.ShouldBindJSON(&param)
	fmt.Printf("userLogin: %+v\n", param)
	c.JSON(gin.H{
		"message": "你好，世界",
	}, nil)
}
