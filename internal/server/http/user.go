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
		r.POST("findAll", box.Handler(t.FindAll, box.AccessAll))
		r.POST("updateUserRole", box.Handler(t.UpdateUserRole, box.AccessAll))
	}
}

func (t *UserServer) UserLogin(c *box.Context) {
	var param model.UserLoginParam
	c.ShouldBindJSON(&param)
	fmt.Printf("userLogin: %+v\n", param)

	data, err := t.service.User.UserLogin(param, c.Req.Passport)
	c.JSON(data, err)
}

func (t *UserServer) FindAll(c *box.Context) {
	var param model.UserFindAllReq
	c.ShouldBindJSON(&param)
	fmt.Printf("FindAll: %+v\n", param)

	data, err := t.service.User.FindAll(param)
	c.JSON(data, err)
}

func (t *UserServer) UpdateUserRole(c *box.Context) {
	var param model.User
	c.ShouldBindJSON(&param)
	fmt.Printf("UpdateUserRole: %+v\n", param)

	err := t.service.User.UpdateUserRole(param)
	c.JSON(nil, err)
}
