package http

import (
	"shippo-server/internal/model"
	"shippo-server/utils/box"
)

type PermissionPolicyServer struct {
	*Server
	router *box.RouterGroup
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
	t.router.POST("create", t.PermissionPolicyCreate)
	t.router.POST("del", t.PermissionPolicyDel)
	t.router.POST("update", t.PermissionPolicyUpdate)
	t.router.POST("updateAccess", t.PermissionAssociationUpdate)
	t.router.POST("findAllExtStatus", t.PermissionPolicyFindAllExtStatus)
	t.router.POST("findAll", t.PermissionPolicyFindAll)
	t.router.POST("find", t.PermissionPolicyFind)
}

func (t *PermissionPolicyServer) PermissionPolicyCreate(c *box.Context) {
	var param model.PermissionPolicy
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}
	err := t.service.PermissionPolicy.PermissionPolicyCreate(param)
	c.JSON(nil, err)
}

func (t *PermissionPolicyServer) PermissionPolicyDel(c *box.Context) {
	var param model.PermissionPolicy
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}
	err := t.service.PermissionPolicy.PermissionPolicyDel(param)
	c.JSON(nil, err)
}

func (t *PermissionPolicyServer) PermissionPolicyUpdate(c *box.Context) {
	var param model.PermissionPolicy
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}
	err := t.service.PermissionPolicy.PermissionPolicyUpdate(param)
	c.JSON(nil, err)
}

func (t *PermissionPolicyServer) PermissionPolicyFindAllExtStatus(c *box.Context) {
	var param model.PermissionPolicy
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}
	data, err := t.service.PermissionPolicy.PermissionPolicyFindAllExtStatus(param.ID)
	c.JSON(data, err)
}

func (t *PermissionPolicyServer) PermissionPolicyFindAll(c *box.Context) {
	data, err := t.service.PermissionPolicy.PermissionPolicyFindAll()
	c.JSON(data, err)
}

func (t *PermissionPolicyServer) PermissionPolicyFind(c *box.Context) {
	var param model.PermissionPolicy
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}
	data, err := t.service.PermissionPolicy.PermissionPolicyFind(param)
	c.JSON(data, err)
}

// 更新权限策略所拥有的访问规则
func (t *PermissionPolicyServer) PermissionAssociationUpdate(c *box.Context) {
	var param = new(struct {
		Id     uint   `json:"id"`
		Access []uint `json:"access"`
	})
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}
	err := t.service.PermissionPolicy.PermissionAssociationUpdate(param.Id, param.Access)
	c.JSON(nil, err)
}
