package http

import (
	"shippo-server/internal/model"
	"shippo-server/utils/box"
)

type WxOffiaccountServer struct {
	*Server
	router *box.RouterGroup
}

func NewWxOffiaccountServer(server *Server) *WxOffiaccountServer {
	var s = &WxOffiaccountServer{
		Server: server,
		router: server.router.Group("wxOffiaccount"),
	}
	s.initRouter()
	return s
}

func (t *WxOffiaccountServer) initRouter() {
	t.router.POST("findAll", t.FindAll)
	t.router.POST("find", t.Find)
}

//FindAll 查询所有公众号
func (t *WxOffiaccountServer) FindAll(c *box.Context) {
	var param model.WxOffiaccount
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}

	r, err := t.service.WxOffiaccount.FindAll()
	c.JSON(r, err)
}

// FindByUsername 根据username查询公众号
func (t *WxOffiaccountServer) FindByUsername(c *box.Context) {
	var param model.WxOffiaccount
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}

	r, err := t.service.WxOffiaccount.FindByUsername(param.Username)
	c.JSON(r, err)
}

// Find 根据id查询公众号
func (t *WxOffiaccountServer) Find(c *box.Context) {
	var param model.WxOffiaccount
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}

	r, err := t.service.WxOffiaccount.Find(param.ID)
	c.JSON(r, err)
}
