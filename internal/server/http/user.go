package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shippo-server/internal/model"
	"shippo-server/utils/box"
)

func initUserRouter(Router *gin.RouterGroup) {
	r := Router.Group("user")
	{
		r.POST("login", box.Handler(userLogin))
	}
}

func userLogin(c *box.Context) {
	var param model.UserLoginParam
	c.ShouldBindJSON(&param)
	fmt.Printf("userLogin: %+v\n", param)

	data, err := svc.UserLogin(c, param, c.Req.Passport)
	c.JSON(data, err)
}
