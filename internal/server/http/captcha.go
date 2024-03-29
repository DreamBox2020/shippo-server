package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"shippo-server/utils/box"
	"shippo-server/utils/config"
)

type CaptchaServer struct {
	*Server
	router *box.RouterGroup
}

func NewCaptchaServer(server *Server) *CaptchaServer {
	var s = &CaptchaServer{
		Server: server,
		router: server.router.Group("captcha"),
	}
	s.initRouter()
	return s
}

func (t *CaptchaServer) initRouter() {
	t.router.POST("send", t.CaptchaSend)
	t.router.GinGroup.Any("serverInfo", t.ServerInfo)
}

func (t *CaptchaServer) CaptchaSend(c *box.Context) {
	var param = new(struct {
		Phone   string `json:"phone"`
		Email   string `json:"email"`
		Channel string `json:"channel"`
	})
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}
	fmt.Printf("c.ShouldBindJSON->param:%+v\n", param)

	if param.Channel == "" {
		param.Channel = "login"
	}

	if param.Phone != "" {
		err := t.service.Captcha.CaptchaSmsSend(param.Phone, c.Req.Passport, param.Channel)
		c.JSON(nil, err)
	} else {
		err := t.service.Captcha.CaptchaEmailSend(param.Email, c.Req.Passport, param.Channel)
		c.JSON(nil, err)
	}
}

func (t *CaptchaServer) ServerInfo(c *gin.Context) {

	c.SetCookie("__ServerInfo", "ServerInfo", 60*60*24*30,
		"/", config.Server.CookieDomain, false, true)

	body, _ := ioutil.ReadAll(c.Request.Body)

	c.JSON(200, gin.H{
		"RequestURI":     c.Request.RequestURI,
		"Host":           c.Request.Host,
		"URL.Host":       c.Request.URL.Host,
		"URL.Path":       c.Request.URL.Path,
		"URL.Fragment":   c.Request.URL.Fragment,
		"URL.Opaque":     c.Request.URL.Opaque,
		"URL.RawPath":    c.Request.URL.RawPath,
		"URL.RawQuery":   c.Request.URL.RawQuery,
		"URL.Scheme":     c.Request.URL.Scheme,
		"URL.RequestURI": c.Request.RequestURI,
		"Method":         c.Request.Method,
		"Header":         c.Request.Header,
		"Body":           string(body),
	})
}
