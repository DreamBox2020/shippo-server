package http

import (
	"shippo-server/internal/model"
	"shippo-server/utils/box"
)

type PermissionAccessServer struct {
	*Server
	router *box.RouterGroup
}

func NewPermissionAccessServer(server *Server) *PermissionAccessServer {
	var s = &PermissionAccessServer{
		Server: server,
		router: server.router.Group("permissionAccess"),
	}
	s.initRouter()
	return s
}

func (t *PermissionAccessServer) initRouter() {
	t.router.POST("create", t.PermissionAccessCreate)
	t.router.POST("del", t.PermissionAccessDel)
	t.router.POST("update", t.PermissionAccessUpdate)
	t.router.POST("findAllExtStatus", t.PermissionAccessFindAllExtStatus)
	t.router.POST("findAll", t.PermissionAccessFindAll)
	t.router.POST("find", t.PermissionAccessFind)
}

func (t *PermissionAccessServer) PermissionAccessCreate(c *box.Context) {
	var p model.PermissionAccess
	c.ShouldBindJSON(&p)
	err := t.service.PermissionAccess.PermissionAccessCreate(p)
	c.JSON(nil, err)
}

func (t *PermissionAccessServer) PermissionAccessDel(c *box.Context) {
	var p model.PermissionAccess
	c.ShouldBindJSON(&p)
	err := t.service.PermissionAccess.PermissionAccessDel(p)
	c.JSON(p, err)
}

func (t *PermissionAccessServer) PermissionAccessUpdate(c *box.Context) {
	var p model.PermissionAccess
	c.ShouldBindJSON(&p)
	err := t.service.PermissionAccess.PermissionAccessUpdate(p)
	c.JSON(nil, err)
}

func (t *PermissionAccessServer) PermissionAccessFindAllExtStatus(c *box.Context) {
	var p model.PermissionAccess
	c.ShouldBindJSON(&p)
	data, err := t.service.PermissionAccess.PermissionAccessFindAllExtStatus(p.ID)
	c.JSON(data, err)
}

func (t *PermissionAccessServer) PermissionAccessFindAll(c *box.Context) {
	data, err := t.service.PermissionAccess.PermissionAccessFindAll()
	c.JSON(data, err)
}

func (t *PermissionAccessServer) PermissionAccessFind(c *box.Context) {
	var p model.PermissionAccess
	c.ShouldBindJSON(&p)
	data, err := t.service.PermissionAccess.PermissionAccessFind(p)
	c.JSON(data, err)
}
