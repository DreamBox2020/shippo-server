package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shippo-server/internal/model"
	"shippo-server/utils/box"
)

type UserServer struct {
	*Server
}

func NewUserServer(s *Server) *UserServer {
	return &UserServer{s}
}

func (t *UserServer) InitRouter(Router *gin.RouterGroup) {
	var h = box.NewBoxHandler(&t)
	r := Router.Group("user")
	{
		r.POST("login", h.H(t.UserLogin, box.AccessAll))
	}
}

func (t *UserServer) UserLogin(c *box.Context) {
	var param model.UserLoginParam
	c.ShouldBindJSON(&param)
	fmt.Printf("userLogin: %+v\n", param)

	data, err := t.service.User.UserLogin(c, param, c.Req.Passport)
	c.JSON(data, err)
}
