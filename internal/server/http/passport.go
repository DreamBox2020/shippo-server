package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shippo-server/utils/box"
	"shippo-server/utils/ecode"
)

func initPassportRouter(Router *gin.RouterGroup) {
	r := Router.Group("passport")
	{
		r.POST("create", box.Handler(passportCreate, box.AccessAll))
	}
}

func passportCreate(c *box.Context) {
	data, err := svc.PassportCreate(c, c.Req.Passport, c.Ctx.ClientIP())
	c.JSON(data, err)
}

func passportGet(c *box.Context) {
	p, err := svc.PassportGet(c, c.Req.Passport, c.Ctx.ClientIP())
	if err != nil {
		fmt.Printf("http->passportGet:%+v\n", err)
		c.JSON(nil, err)
		return
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
