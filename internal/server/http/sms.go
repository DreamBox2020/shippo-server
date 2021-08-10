package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shippo-server/utils/box"
)

func initSmsRouter(Router *gin.RouterGroup) {
	r := Router.Group("sms")
	{
		r.POST("send", box.Handler(smsSend, box.AccessAll))
	}
}

func smsSend(c *box.Context) {
	var param = new(struct {
		Phone string `json:"phone"`
	})
	c.ShouldBindJSON(&param)
	fmt.Printf("smsSend: %+v\n", param)

	err := svc.SmsSend(c, param.Phone, c.Req.Passport)
	c.JSON(nil, err)
}
