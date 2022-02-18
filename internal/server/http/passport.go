package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
	var h = box.NewBoxHandler(&t)

	r := Router.Group("passport")
	{
		r.POST("create", h.H(t.PassportCreate, box.AccessAll))
	}
}

func (t *PassportServer) PassportCreate(c *box.Context) {
	data, err := t.service.Passport.PassportCreate(c, c.Req.Passport, c.Ctx.ClientIP())
	c.JSON(data, err)
}

func (t *PassportServer) PassportGet(c *box.Context) {
	p, err := t.service.Passport.PassportGet(c, c.Req.Passport, c.Ctx.ClientIP())
	if err != nil {
		fmt.Printf("http->passportGet:%+v\n", err)
	}
	fmt.Printf("http->passportGet:%+v\n", p)
	c.Passport = &p

	// 如果需要登录权限，但是并没有登录。
	if c.Access == box.AccessLoginOK {
		if c.Passport.UserId == 0 {
			c.JSON(nil, ecode.NoLogin)
			return
		}
	}

	c.Next()
}
