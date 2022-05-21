package http

import (
	"github.com/gin-gonic/gin"
	"shippo-server/internal/model"
	"shippo-server/utils/box"
)

type PermissionAccessServer struct {
	*Server
}

func NewPermissionAccessServer(s *Server) *PermissionAccessServer {
	return &PermissionAccessServer{s}
}

func (t *PermissionAccessServer) InitRouter(Router *gin.RouterGroup) {
	r := Router.Group("permissionAccess")
	{
		r.POST("create", box.Handler(t.PermissionAccessCreate))
		r.POST("del", box.Handler(t.PermissionAccessDel))
		r.POST("update", box.Handler(t.PermissionAccessUpdate))
		r.POST("findAllExtStatus", box.Handler(t.PermissionAccessFindAllExtStatus))
		r.POST("findAll", box.Handler(t.PermissionAccessFindAll))
		r.POST("find", box.Handler(t.PermissionAccessFind))
	}
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
