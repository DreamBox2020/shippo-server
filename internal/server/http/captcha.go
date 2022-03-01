package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shippo-server/utils/box"
)

type CaptchaServer struct {
	*Server
}

func NewCaptchaServer(s *Server) *CaptchaServer {
	return &CaptchaServer{s}
}

func (t *CaptchaServer) InitRouter(Router *gin.RouterGroup) {
	r := Router.Group("captcha")
	{
		r.POST("send", box.Handler(t.CaptchaSend, box.AccessAll))
	}
}

func (t *CaptchaServer) CaptchaSend(c *box.Context) {
	var param = new(struct {
		Phone string `json:"phone"`
		Email string `json:"email"`
	})
	c.ShouldBindJSON(&param)
	fmt.Printf("captchaSend: %+v\n", param)

	if param.Phone != "" {
		err := t.service.Captcha.CaptchaSmsSend(param.Phone, c.Req.Passport)
		c.JSON(nil, err)
	} else {
		err := t.service.Captcha.CaptchaEmailSend(param.Email, c.Req.Passport)
		c.JSON(nil, err)
	}
}
