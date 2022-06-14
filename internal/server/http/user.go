package http

import (
	"fmt"
	"shippo-server/internal/model"
	"shippo-server/utils"
	"shippo-server/utils/box"
)

type UserServer struct {
	*Server
	router *box.RouterGroup
}

func NewUserServer(server *Server) *UserServer {
	var s = &UserServer{
		Server: server,
		router: server.router.Group("user"),
	}
	s.initRouter()
	return s
}

func (t *UserServer) initRouter() {
	t.router.POST("login", t.UserLogin)
	t.router.POST("logout", t.UserLogout)
	t.router.POST("findAll", t.FindAll)
	t.router.POST("updateUserRole", t.UpdateUserRole)
}

func (t *UserServer) UserLogout(c *box.Context) {

	access, err := t.service.PermissionPolicy.FindPermissionAccessByPolicyName("SysBase")

	if err != nil {
		c.JSON(nil, err)
		return
	}

	passport, err := t.service.Passport.CreateNoLoginPassport(*c.Passport)
	if err != nil {
		c.JSON(nil, err)
		return
	}

	var data model.PassportCreateResult
	data.Passport = passport.Token
	data.Uid = 0
	data.Access = access

	data.User = model.User{}

	c.JSON(data, err)
}

func (t *UserServer) UserLogin(c *box.Context) {
	var param model.UserLoginParam
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}
	fmt.Printf("c.ShouldBindJSON->param:%+v\n", param)

	user, err := t.service.User.UserLogin(param, *c.Passport)
	if err != nil {
		c.JSON(nil, err)
		return
	}

	access, err := t.service.Role.RoleFindPermissionAccess(user.Role)
	if err != nil {
		c.JSON(nil, err)
		return
	}

	var data model.PassportCreateResult
	data.Passport = c.Passport.Token
	data.Uid = user.ID
	data.Access = access

	user.Email = utils.QQEmailMasking(user.Email)
	user.Phone = utils.PhoneMasking(user.Phone)
	data.User = user

	c.JSON(data, err)
}

func (t *UserServer) FindAll(c *box.Context) {
	var param model.UserFindAllReq
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}
	fmt.Printf("c.ShouldBindJSON->param:%+v\n", param)

	data, err := t.service.User.FindAll(param)
	c.JSON(data, err)
}

func (t *UserServer) UpdateUserRole(c *box.Context) {
	var param model.User
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}
	fmt.Printf("c.ShouldBindJSON->param:%+v\n", param)

	err := t.service.User.UpdateUserRole(param)
	c.JSON(nil, err)
}
