package http

import (
	"shippo-server/internal/model"
	"shippo-server/utils/box"
)

type WxCommentServer struct {
	*Server
	router *box.RouterGroup
}

func NewWxCommentServer(server *Server) *WxCommentServer {
	var s = &WxCommentServer{
		Server: server,
		router: server.router.Group("wxComment"),
	}
	s.initRouter()
	return s
}

func (t *WxCommentServer) initRouter() {
	t.router.POST("create", t.Create)
	t.router.POST("reply", t.Reply)
	t.router.POST("admin/reply", t.AdminReply)
	t.router.POST("delete", t.Delete)
	t.router.POST("updateElected", t.UpdateElected)
	t.router.POST("updateTop", t.UpdateTop)
	t.router.POST("findByWxPassportAndOffiaccount", t.FindByWxPassportAndOffiaccount)
	t.router.POST("findByWxPassportAndOffiaccountAndElected", t.FindByWxPassportAndOffiaccountAndElected)
	t.router.POST("findByWxPassportAndArticle", t.FindByWxPassportAndArticle)
	t.router.POST("findByArticle", t.FindByArticle)
	t.router.POST("admin/findByArticle", t.AdminFindByArticle)
}

func (t *WxCommentServer) Create(c *box.Context) {
	var param *model.WxComment
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}

	param.WxPassportId = c.User.WxPassportId

	_, err := t.service.WxComment.Create(param)
	c.JSON(nil, err)
}

func (t *WxCommentServer) Reply(c *box.Context) {
	var param *model.WxComment
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}

	param.WxPassportId = c.User.WxPassportId

	r, err := t.service.WxComment.Reply(param)
	c.JSON(r, err)
}

func (t *WxCommentServer) AdminReply(c *box.Context) {
	var param *model.WxComment
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}

	param.WxPassportId = c.User.WxPassportId

	r, err := t.service.WxComment.AdminReply(param)
	c.JSON(r, err)
}

func (t *WxCommentServer) Delete(c *box.Context) {
	var param *model.WxComment
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}

	param.WxPassportId = c.User.WxPassportId

	err := t.service.WxComment.Delete(param)
	c.JSON(nil, err)
}

func (t *WxCommentServer) UpdateElected(c *box.Context) {
	var param *model.WxComment
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}

	param.WxPassportId = c.User.WxPassportId

	err := t.service.WxComment.UpdateElected(param)
	c.JSON(nil, err)
}

func (t *WxCommentServer) UpdateTop(c *box.Context) {
	var param *model.WxComment
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}

	param.WxPassportId = c.User.WxPassportId

	err := t.service.WxComment.UpdateTop(param)
	c.JSON(nil, err)
}

func (t *WxCommentServer) FindByWxPassportAndOffiaccount(c *box.Context) {
	var param *model.WxArticle
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}

	param.WxPassportId = c.User.WxPassportId

	r, err := t.service.WxComment.FindByWxPassportAndOffiaccount(param)
	c.JSON(r, err)
}

func (t *WxCommentServer) FindByWxPassportAndOffiaccountAndElected(c *box.Context) {
	var param *model.WxArticle
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}

	param.WxPassportId = c.User.WxPassportId

	r, err := t.service.WxComment.FindByWxPassportAndOffiaccountAndElected(param)
	c.JSON(r, err)
}

func (t *WxCommentServer) FindByWxPassportAndArticle(c *box.Context) {
	var param *model.WxComment
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}

	param.WxPassportId = c.User.WxPassportId

	r, err := t.service.WxComment.FindByWxPassportAndArticle(param)
	c.JSON(r, err)
}

func (t *WxCommentServer) FindByArticle(c *box.Context) {
	var param *model.WxComment
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}

	r, err := t.service.WxComment.FindByArticle(param)
	c.JSON(r, err)
}

func (t *WxCommentServer) AdminFindByArticle(c *box.Context) {
	var param *model.WxComment
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}
	
	r, err := t.service.WxComment.AdminFindByArticle(param)
	c.JSON(r, err)
}
