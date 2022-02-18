package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shippo-server/utils/box"
)

type AdminUserServer struct {
	*Server
}

func NewAdminUserServer(s *Server) *AdminUserServer {
	return &AdminUserServer{s}
}

func (t *AdminUserServer) InitRouter(Router *gin.RouterGroup) {
	var h = box.NewBoxHandler(&t)

	r := Router.Group("admin/user")
	{
		r.POST("create", h.H(t.UserCreateEmail, box.AccessLoginOK))
	}
}

func (t *AdminUserServer) UserCreateEmail(c *box.Context) {
	var param = new(struct {
		Email string `json:"email"`
	})
	c.ShouldBindJSON(&param)
	fmt.Printf("userCreateEmail: %+v\n", param)

	_, err := t.service.AdminUser.AdminUserCreateEmail(c, param.Email)
	c.JSON(nil, err)
}
