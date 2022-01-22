package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shippo-server/utils/box"
)

func initCaptchaRouter(Router *gin.RouterGroup) {
	r := Router.Group("captcha")
	{
		r.POST("send", box.Handler(captchaSend, box.AccessAll))
	}
}

func captchaSend(c *box.Context) {
	var param = new(struct {
		Phone string `json:"phone"`
		Email string `json:"email"`
	})
	c.ShouldBindJSON(&param)
	fmt.Printf("captchaSend: %+v\n", param)

	if param.Phone != "" {
		err := svc.CaptchaSmsSend(c, param.Phone, c.Req.Passport)
		c.JSON(nil, err)
	} else {
		err := svc.CaptchaEmailSend(c, param.Email, c.Req.Passport)
		c.JSON(nil, err)
	}
}
