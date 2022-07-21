package http

import (
	"fmt"
	"shippo-server/internal/model"
	"shippo-server/utils/box"
	"shippo-server/utils/ecode"
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
	t.router.POST("del", t.Delete)
	t.router.POST("updateElected", t.UpdateElected)
	t.router.POST("updateTop", t.UpdateTop)
	t.router.POST("findByWxPassportAndOffiaccount", t.FindByWxPassportAndOffiaccount)
	t.router.POST("findByWxPassportAndOffiaccountAndElected", t.FindByWxPassportAndOffiaccountAndElected)
	t.router.POST("findByWxPassportAndArticle", t.FindByWxPassportAndArticle)
	t.router.POST("findByArticle", t.FindByArticle)
	t.router.POST("admin/findByArticle", t.AdminFindByArticle)
}

// Create 创建一条评论
func (t *WxCommentServer) Create(c *box.Context) {

	if c.User.WxPassportId == 0 {
		c.JSON(nil, ecode.WxPassportIsNull)
		return
	}

	var param *model.WxComment
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}
	fmt.Printf("c.ShouldBindJSON->param:%+v\n", param)

	param.WxPassportId = c.User.WxPassportId

	r, err := t.service.WxComment.Create(param)
	c.JSON(r, err)
}

// Reply 回复一条评论
func (t *WxCommentServer) Reply(c *box.Context) {

	if c.User.WxPassportId == 0 {
		c.JSON(nil, ecode.WxPassportIsNull)
		return
	}

	var param *model.WxComment
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}
	fmt.Printf("c.ShouldBindJSON->param:%+v\n", param)

	param.WxPassportId = c.User.WxPassportId

	r, err := t.service.WxComment.Reply(param)
	c.JSON(r, err)
}

// AdminReply 管理员回复一条评论
func (t *WxCommentServer) AdminReply(c *box.Context) {

	if c.User.WxPassportId == 0 {
		c.JSON(nil, ecode.WxPassportIsNull)
		return
	}

	var param *model.WxComment
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}
	fmt.Printf("c.ShouldBindJSON->param:%+v\n", param)

	param.WxPassportId = c.User.WxPassportId

	r, err := t.service.WxComment.AdminReply(param)
	c.JSON(r, err)
}

// Delete 删除一条评论
func (t *WxCommentServer) Delete(c *box.Context) {

	if c.User.WxPassportId == 0 {
		c.JSON(nil, ecode.WxPassportIsNull)
		return
	}

	var param *model.WxComment
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}
	fmt.Printf("c.ShouldBindJSON->param:%+v\n", param)

	param.WxPassportId = c.User.WxPassportId

	err := t.service.WxComment.Delete(param)
	c.JSON(nil, err)
}

// UpdateElected 更新评论精选状态
func (t *WxCommentServer) UpdateElected(c *box.Context) {

	if c.User.WxPassportId == 0 {
		c.JSON(nil, ecode.WxPassportIsNull)
		return
	}

	var param *model.WxComment
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}
	fmt.Printf("c.ShouldBindJSON->param:%+v\n", param)

	param.WxPassportId = c.User.WxPassportId

	err := t.service.WxComment.UpdateElected(param)
	c.JSON(nil, err)
}

// UpdateTop 更新评论置顶状态
func (t *WxCommentServer) UpdateTop(c *box.Context) {

	if c.User.WxPassportId == 0 {
		c.JSON(nil, ecode.WxPassportIsNull)
		return
	}

	var param *model.WxComment
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}
	fmt.Printf("c.ShouldBindJSON->param:%+v\n", param)

	param.WxPassportId = c.User.WxPassportId

	err := t.service.WxComment.UpdateTop(param)
	c.JSON(nil, err)
}

func (t *WxCommentServer) FindByWxPassportAndOffiaccount(c *box.Context) {
	var param *model.WxArticle
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}
	fmt.Printf("c.ShouldBindJSON->param:%+v\n", param)

	param.WxPassportId = c.User.WxPassportId

	r, err := t.service.WxComment.FindByWxPassportAndOffiaccount(param)
	c.JSON(r, err)
}

func (t *WxCommentServer) FindByWxPassportAndOffiaccountAndElected(c *box.Context) {
	var param *model.WxArticle
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}
	fmt.Printf("c.ShouldBindJSON->param:%+v\n", param)

	param.WxPassportId = c.User.WxPassportId

	r, err := t.service.WxComment.FindByWxPassportAndOffiaccountAndElected(param)
	c.JSON(r, err)
}

// FindByWxPassportAndArticle 查询某用户在某文章的全部评论
func (t *WxCommentServer) FindByWxPassportAndArticle(c *box.Context) {
	var param *model.WxComment
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}
	fmt.Printf("c.ShouldBindJSON->param:%+v\n", param)

	param.WxPassportId = c.User.WxPassportId

	r, err := t.service.WxComment.FindByWxPassportAndArticle(param)
	c.JSON(r, err)
}

// FindByArticle 查询某文章的全部精选评论
func (t *WxCommentServer) FindByArticle(c *box.Context) {
	var param *model.WxComment
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}
	fmt.Printf("c.ShouldBindJSON->param:%+v\n", param)

	r, err := t.service.WxComment.FindByArticle(param)
	c.JSON(r, err)
}

// AdminFindByArticle 查询某文章的全部评论，包含未精选
func (t *WxCommentServer) AdminFindByArticle(c *box.Context) {
	var param *model.WxComment
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}
	fmt.Printf("c.ShouldBindJSON->param:%+v\n", param)

	r, err := t.service.WxComment.AdminFindByArticle(param)
	c.JSON(r, err)
}
