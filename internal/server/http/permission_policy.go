package http

import (
	"github.com/gin-gonic/gin"
	"shippo-server/internal/model"
	"shippo-server/utils/box"
)

type PermissionPolicyServer struct {
	*Server
	router *gin.RouterGroup
}

func NewPermissionPolicyServer(server *Server) *PermissionPolicyServer {
	var s = &PermissionPolicyServer{
		Server: server,
		router: server.router.Group("permissionPolicy"),
	}
	s.initRouter()
	return s
}

func (t *PermissionPolicyServer) initRouter() {
	t.router.POST("create", box.Handler(t.PermissionPolicyCreate))
	t.router.POST("del", box.Handler(t.PermissionPolicyDel))
	t.router.POST("update", box.Handler(t.PermissionPolicyUpdate))
	t.router.POST("updateAccess", box.Handler(t.PermissionAssociationUpdate))
	t.router.POST("findAllExtStatus", box.Handler(t.PermissionPolicyFindAllExtStatus))
	t.router.POST("findAll", box.Handler(t.PermissionPolicyFindAll))
	t.router.POST("find", box.Handler(t.PermissionPolicyFind))
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

func (t *PermissionPolicyServer) PermissionPolicyFindAllExtStatus(c *box.Context) {
	var p model.PermissionPolicy
	c.ShouldBindJSON(&p)
	data, err := t.service.PermissionPolicy.PermissionPolicyFindAllExtStatus(p.ID)
	c.JSON(data, err)
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

// 更新权限策略所拥有的访问规则
func (t *PermissionPolicyServer) PermissionAssociationUpdate(c *box.Context) {
	var param = new(struct {
		Id     uint   `json:"id"`
		Access []uint `json:"access"`
	})
	c.ShouldBindJSON(&param)
	err := t.service.PermissionPolicy.PermissionAssociationUpdate(param.Id, param.Access)
	c.JSON(nil, err)
}
