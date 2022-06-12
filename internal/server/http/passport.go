package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"shippo-server/internal/model"
	"shippo-server/utils/box"
	"shippo-server/utils/config"
	"shippo-server/utils/ecode"
	"strings"
)

type PassportServer struct {
	*Server
	router *box.RouterGroup
}

func NewPassportServer(server *Server) *PassportServer {
	var s = &PassportServer{
		Server: server,
		router: server.router.Group("passport"),
	}
	s.initRouter()
	return s
}

func (t *PassportServer) initRouter() {
	t.router.POST("create", t.PassportCreate)
	t.router.GinGroup.POST("createDev", t.CreateDev)
}

func (t *PassportServer) PassportCreate(c *box.Context) {

	var param = new(struct {
		Code string `json:"code"`
	})

	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}
	fmt.Printf("c.ShouldBindJSON->param:%+v\n", param)

	var passport model.Passport
	var err error = nil

	if param.Code != "" {
		passport, err = t.service.Passport.WxCreate(*c.Passport, param.Code)
		if err != nil {
			c.JSON(nil, err)
			return
		}
	} else {
		passport, err = t.service.Passport.PassportCreate(*c.Passport)
		if err != nil {
			c.JSON(nil, err)
			return
		}
	}

	var domain string
	if !config.IsLocal() {
		domain = config.Server.CookieDomain
	}
	c.SetCookie("__PASSPORT", passport.Token, 60*60*24*30,
		"/", domain, false, true)

	var data model.PassportCreateResult
	var access []model.PermissionAccess

	if passport.IsLogin() {
		// 根据用户角色查询对应权限信息
		access, err = t.service.Role.RoleFindPermissionAccess(c.User.Role)
	} else {
		// 如果当前没有登录，查询基础权限信息
		access, err = t.service.PermissionPolicy.FindPermissionAccessByPolicyName("SysBase")
	}

	if err != nil {
		c.JSON(nil, err)
		return
	}

	data.Passport = passport.Token
	data.Uid = passport.UserId
	data.Access = access

	c.JSON(data, err)
}

func (t *PassportServer) CreateDev(c *gin.Context) {

	if !config.IsLocal() {
		c.String(http.StatusNotFound, "404 page not found")
		return
	}

	var param = new(struct {
		Uid uint `json:"uid"`
	})

	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}
	fmt.Printf("c.ShouldBindJSON->param:%+v\n", param)

	data, _ := t.service.Passport.CreateLoginPassport(model.Passport{
		UserId: param.Uid,
		Ip:     c.ClientIP(),
		Ua:     c.GetHeader("User-Agent"),
		Client: 0,
	})

	c.String(http.StatusOK, data.Token)

}

func (t *PassportServer) PassportGet(c *box.Context) {
	fmt.Printf("http->passportGet->ip:%+v\n", c.ClientIP())
	fmt.Printf("http->passportGet->ua:%+v\n", c.GetHeader("User-Agent"))

	if c.Req.Passport != "" {
		p, err := t.service.Passport.PassportGet(c.Req.Passport)
		if err != nil {
			fmt.Printf("http->passportGet->err:%+v\n", err)
			c.JSON(nil, ecode.ServerErr)
			c.Abort()
			return
		}
		fmt.Printf("http->passportGet:%+v\n", p)
		p.Ip = c.ClientIP()
		p.Ua = c.GetHeader("User-Agent")
		c.Passport = &p
	} else {
		c.Passport = &model.Passport{
			Ip:     c.ClientIP(),
			Ua:     c.GetHeader("User-Agent"),
			Client: 0,
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
		r, err := t.service.PermissionPolicy.FindPermissionAccessByPolicyNameAndType(
			"SysBase", "action")
		if err != nil {
			c.JSON(nil, ecode.ServerErr)
			c.Abort()
			return
		}
		list = r

	}

	fmt.Printf("http->Auth->list:%+v\n", list)

	var tag = false
	reg, _ := regexp.Compile("^/v1")
	path := reg.ReplaceAllString(c.Request.URL.Path, "")
	var key1 = strings.ToLower(c.Request.Method + ":" + path)
	for _, access := range list {
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
		if c.Passport.IsLogin() {
			c.JSON(nil, ecode.AccessDenied)
		} else {
			c.JSON(nil, ecode.NoLogin)
		}
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
