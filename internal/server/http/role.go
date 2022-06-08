package http

import (
	"fmt"
	"shippo-server/internal/model"
	"shippo-server/utils/box"
)

type RoleServer struct {
	*Server
	router *box.RouterGroup
}

func NewRoleServer(server *Server) *RoleServer {
	var s = &RoleServer{
		Server: server,
		router: server.router.Group("role"),
	}
	s.initRouter()
	return s
}

func (t *RoleServer) initRouter() {
	t.router.POST("create", t.RoleCreate)
	t.router.POST("del", t.RoleDel)
	t.router.POST("update", t.RoleUpdate)
	t.router.POST("updatePolicies", t.RoleAssociationUpdate)
	t.router.POST("findAll", t.RoleFindAll)
	t.router.POST("findPolicies", t.FindPolicies)
	t.router.POST("find", t.RoleFind)
}

// 增加⻆⾊
func (t *RoleServer) RoleCreate(c *box.Context) {
	var r model.Role
	c.ShouldBindJSON(&r)
	err := t.service.Role.RoleCreate(r)
	c.JSON(nil, err)
}

// 删除⻆⾊
func (t *RoleServer) RoleDel(c *box.Context) {
	var param model.Role
	c.ShouldBindJSON(&param)
	err := t.service.Role.RoleDel(param)
	c.JSON(nil, err)
}

// 更新⻆⾊名称和备注
func (t *RoleServer) RoleUpdate(c *box.Context) {
	var param model.Role
	c.ShouldBindJSON(&param)
	fmt.Printf("RoleUpdate: %+v\n", param)

	err := t.service.Role.RoleUpdate(param)
	c.JSON(nil, err)
}

// 更新角色所拥有的权限策略
func (t *RoleServer) RoleAssociationUpdate(c *box.Context) {
	var param = new(struct {
		Id       uint   `json:"id"`
		Policies []uint `json:"policies"`
	})
	c.ShouldBindJSON(&param)
	err := t.service.Role.RoleAssociationUpdate(param.Id, param.Policies)
	c.JSON(nil, err)
}

// 查询全部⻆⾊
func (t *RoleServer) RoleFindAll(c *box.Context) {
	data, err := t.service.Role.RoleFindAll()
	c.JSON(data, err)
}

// 查询某个⻆⾊所拥有的权限策略
func (t *RoleServer) FindPolicies(c *box.Context) {
	var param model.Role
	c.ShouldBindJSON(&param)
	data, err := t.service.Role.RoleFindPermissionPolicy(param.ID)
	c.JSON(data, err)
}

// 查询某个⻆⾊所拥有的访问规则
func (t *RoleServer) RoleFind(c *box.Context) {
	var param model.Role
	c.ShouldBindJSON(&param)
	data, err := t.service.Role.RoleFindPermissionAccess(param.ID)
	c.JSON(data, err)
}
