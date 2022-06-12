package http

import (
	"fmt"
	"shippo-server/internal/model"
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
	t.router.POST("findAll", t.FindAll)
	t.router.POST("updateUserRole", t.UpdateUserRole)
}

func (t *UserServer) UserLogin(c *box.Context) {
	var param model.UserLoginParam
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}
	fmt.Printf("c.ShouldBindJSON->param:%+v\n", param)

	user, err := t.service.User.UserLogin(param, *c.Passport)
	if err != nil {
		return
	}

	access, err := t.service.Role.RoleFindPermissionAccess(user.Role)
	if err != nil {
		return
	}

	data := make(map[string]interface{})
	data["access"] = access
	data["passport"] = c.Passport.Token
	data["uid"] = user.ID

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
