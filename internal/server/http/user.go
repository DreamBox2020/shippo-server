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
	r := Router.Group("user")
	{
		r.POST("login", box.Handler(t.UserLogin, box.AccessAll))
	}
}

func (t *UserServer) UserLogin(c *box.Context) {
	var param model.UserLoginParam
	c.ShouldBindJSON(&param)
	fmt.Printf("userLogin: %+v\n", param)

	data, err := t.service.User.UserLogin(param, c.Req.Passport)
	c.JSON(data, err)
}
