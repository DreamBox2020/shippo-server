package http

import (
	"fmt"
	"shippo-server/utils/box"
)

type AdminUserServer struct {
	*Server
	router *box.RouterGroup
}

func NewAdminUserServer(server *Server) *AdminUserServer {
	var s = &AdminUserServer{
		Server: server,
		router: server.router.Group("admin/user"),
	}
	s.initRouter()
	return s
}

func (t *AdminUserServer) initRouter() {
	t.router.POST("create", t.UserCreateEmail)
}

func (t *AdminUserServer) UserCreateEmail(c *box.Context) {
	var param = new(struct {
		Email string `json:"email"`
	})
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}
	fmt.Printf("c.ShouldBindJSON->param:%+v\n", param)

	_, err := t.service.AdminUser.AdminUserCreateEmail(param.Email)
	c.JSON(nil, err)
}
