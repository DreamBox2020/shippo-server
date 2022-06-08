package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shippo-server/utils/box"
)

type AdminUserServer struct {
	*Server
	router *gin.RouterGroup
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
	t.router.POST("create", box.Handler(t.UserCreateEmail))
}

func (t *AdminUserServer) UserCreateEmail(c *box.Context) {
	var param = new(struct {
		Email string `json:"email"`
	})
	c.ShouldBindJSON(&param)
	fmt.Printf("userCreateEmail: %+v\n", param)

	_, err := t.service.AdminUser.AdminUserCreateEmail(param.Email)
	c.JSON(nil, err)
}
