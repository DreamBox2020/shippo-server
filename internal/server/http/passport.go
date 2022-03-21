package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"regexp"
	"shippo-server/internal/model"
	"shippo-server/utils/box"
	"shippo-server/utils/ecode"
	"strings"
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

	data, err := t.service.Passport.PassportCreate(*c.Passport)
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
		p, err := t.service.Passport.PassportGet(c.Req.Passport)
		if err != nil {
			fmt.Printf("http->passportGet:%+v\n", err)
			c.JSON(nil, ecode.ServerErr)
			c.Abort()
			return
		}
		fmt.Printf("http->passportGet:%+v\n", p)
		c.Passport = &p
	} else {
		c.Passport = &model.Passport{
			Ip:     c.Ctx.ClientIP(),
			Ua:     c.Ctx.GetHeader("User-Agent"),
			Client: 0,
		}
	}

	// 如果需要登录权限，但是并没有登录。
	if c.Access == box.AccessLoginOK {
		if !c.Passport.IsLogin() {
			c.JSON(nil, ecode.NoLogin)
			c.Abort()
			return
		}
	}

	c.Next()
}

func (t *PassportServer) Auth(c *box.Context) {
	var list []model.PermissionAccess
	// 如果已经登录
	if c.Passport.IsLogin() {
		// 查询用户信息并储存
		u, err := t.service.User.UserFindByUID(c.Passport.UserId)
		if err != nil {
			c.JSON(nil, ecode.ServerErr)
			c.Abort()
			return
		}
		c.User = &u
		fmt.Printf("http->Auth:%+v\n", u)

		// 查询用户角色所拥有的访问规则
		r, err := t.service.Role.RoleFindPermissionAccessByType(u.Role, "action")
		if err != nil {
			c.JSON(nil, ecode.ServerErr)
			c.Abort()
			return
		}
		list = r

	} else {

		// 储存一个空的用户信息
		c.User = &model.User{}

		// 查询系统基本访问策略所拥有的访问规则
		r, err := t.service.PermissionPolicy.FindPermissionAccessByPolicyNameAndType("SysBase", "action")
		if err != nil {
			c.JSON(nil, ecode.ServerErr)
			c.Abort()
			return
		}
		list = r

	}

	fmt.Printf("http->Auth->list:%+v\n", list)

	var tag = false
	for _, access := range list {
		key1 := strings.ToLower(c.Ctx.Request.Method + ":" + c.Ctx.Request.URL.Path)
		key2 := strings.ToLower(access.AccessRule)
		fmt.Printf("http->Auth->key1:%+v\n", key1)
		fmt.Printf("http->Auth->key2:%+v\n", key2)

		tag = KeyMatch2(key1, key2)
		fmt.Printf("http->Auth->KeyMatch2:%+v\n", tag)

		if tag {
			break
		}
	}

	if tag {
		c.Next()
	} else {
		c.JSON(nil, ecode.AccessDenied)
		c.Abort()
		return
	}
}

func KeyMatch2(key1 string, key2 string) bool {
	key2 = strings.Replace(key2, "/*", "/.*", -1)

	re := regexp.MustCompile(`:[^/]+`)
	key2 = re.ReplaceAllString(key2, "$1[^/]+$2")

	return RegexMatch(key1, "^"+key2+"$")
}

func RegexMatch(key1 string, key2 string) bool {
	res, err := regexp.MatchString(key2, key1)
	if err != nil {
		panic(err)
	}
	return res
}
