package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shippo-server/internal/model"
	"shippo-server/utils/box"
	"shippo-server/utils/ecode"
)

type PassportServer struct {
	*Server
}

func NewPassportServer(s *Server) *PassportServer {
	return &PassportServer{s}
}

func (t *PassportServer) InitRouter(Router *gin.RouterGroup) {
	r := Router.Group("passport")
	{
		r.POST("create", box.Handler(t.PassportCreate, box.AccessAll))
	}
}

func (t *PassportServer) PassportCreate(c *box.Context) {
	data, err := t.service.Passport.PassportCreate(c.Req.Passport, c.Ctx.ClientIP())
	if err == nil {
		var domain string
		if c.Ctx.ClientIP() != "127.0.0.1" {
			domain = serverConf.CookieDomain
		}
		c.Ctx.SetCookie("__PASSPORT", data.Passport, 60*60*24*30, "/", domain, false, true)
	}
	c.JSON(data, err)
}

func (t *PassportServer) PassportGet(c *box.Context) {

	if c.Req.Passport != "" {
		p, err := t.service.Passport.PassportGet(c.Req.Passport, c.Ctx.ClientIP())
		if err != nil {
			fmt.Printf("http->passportGet:%+v\n", err)
			c.JSON(nil, ecode.ServerErr)
			return
		}
		fmt.Printf("http->passportGet:%+v\n", p)
		c.Passport = &p
	} else {
		c.Passport = &model.Passport{}
	}

	// 如果需要登录权限，但是并没有登录。
	if c.Access == box.AccessLoginOK {
		if c.Passport.UserId == 0 {
			c.JSON(nil, ecode.NoLogin)
			return
		}
	}

	c.Next()
}
