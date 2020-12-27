package http

import (
	"github.com/gin-gonic/gin"
	"shippo-server/utils/box"
)

func initPassportRouter(Router *gin.RouterGroup) {
	r := Router.Group("passport")
	{
		r.POST("create", box.Handler(passportCreate))
	}
}

func passportCreate(c *box.Context) {
	svc.PassportCreate(c, c.Req.Passport)
	c.JSON(nil, nil)
}
