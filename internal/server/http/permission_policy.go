package http

import (
	"github.com/gin-gonic/gin"
	"shippo-server/internal/model"
	"shippo-server/utils/box"
)

type PermissionPolicyServer struct {
	*Server
}

func NewPermissionPolicyServer(s *Server) *PermissionPolicyServer {
	return &PermissionPolicyServer{s}
}

func (t *PermissionPolicyServer) InitRouter(Router *gin.RouterGroup) {
	r := Router.Group("permissionPolicy")
	{
		r.POST("create", box.Handler(t.PermissionPolicyCreate, box.AccessAll))
		r.POST("del", box.Handler(t.PermissionPolicyDel, box.AccessAll))
		r.POST("update", box.Handler(t.PermissionPolicyUpdate, box.AccessAll))
		r.POST("findAll", box.Handler(t.PermissionPolicyFindAll, box.AccessAll))
		r.POST("find", box.Handler(t.PermissionPolicyFind, box.AccessAll))
	}
}

func (t *PermissionPolicyServer) PermissionPolicyCreate(c *box.Context) {
	var p model.PermissionPolicy
	c.ShouldBindJSON(&p)
	err := t.service.PermissionPolicy.PermissionPolicyCreate(p)
	c.JSON(nil, err)
}

func (t *PermissionPolicyServer) PermissionPolicyDel(c *box.Context) {
	var p model.PermissionPolicy
	c.ShouldBindJSON(&p)
	err := t.service.PermissionPolicy.PermissionPolicyDel(p)
	c.JSON(nil, err)
}

func (t *PermissionPolicyServer) PermissionPolicyUpdate(c *box.Context) {
	var p model.PermissionPolicy
	c.ShouldBindJSON(&p)
	err := t.service.PermissionPolicy.PermissionPolicyUpdate(p)
	c.JSON(nil, err)
}

func (t *PermissionPolicyServer) PermissionPolicyFindAll(c *box.Context) {
	data, err := t.service.PermissionPolicy.PermissionPolicyFindAll()
	c.JSON(data, err)
}

func (t *PermissionPolicyServer) PermissionPolicyFind(c *box.Context) {
	var p model.PermissionPolicy
	c.ShouldBindJSON(&p)
	data, err := t.service.PermissionPolicy.PermissionPolicyFind(p)
	c.JSON(data, err)
}
