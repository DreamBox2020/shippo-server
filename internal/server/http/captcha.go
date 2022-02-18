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
	var h = box.NewBoxHandler(&t)

	r := Router.Group("captcha")
	{
		r.POST("send", h.H(t.CaptchaSend, box.AccessAll))
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
		err := t.service.Captcha.CaptchaSmsSend(c, param.Phone, c.Req.Passport)
		c.JSON(nil, err)
	} else {
		err := t.service.Captcha.CaptchaEmailSend(c, param.Email, c.Req.Passport)
		c.JSON(nil, err)
	}
}
