package http

import (
	"shippo-server/internal/model"
	"shippo-server/utils/box"
	"shippo-server/utils/ecode"
)

type WxArticleServer struct {
	*Server
	router *box.RouterGroup
}

func NewWxArticleServer(server *Server) *WxArticleServer {
	var s = &WxArticleServer{
		Server: server,
		router: server.router.Group("wxArticle"),
	}
	s.initRouter()
	return s
}
func (t *WxArticleServer) initRouter() {
	t.router.POST("create", t.Create)
	t.router.POST("update", t.Update)
	t.router.POST("updateCommentSwitch", t.UpdateCommentSwitch)
	t.router.POST("findByOffiaccount", t.FindByOffiaccount)
	t.router.POST("find", t.Find)
	t.router.POST("findAllByWxPassport", t.FindAllByWxPassport)
}

// Create 新增文章
func (t *WxArticleServer) Create(c *box.Context) {
	if c.User.WxPassportId == 0 {
		c.JSON(nil, ecode.WxPassportIsNull)
		return
	}

	var param model.WxArticle
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}

	param.WxPassportId = c.User.WxPassportId

	_, err := t.service.WxArticle.Create(&param)
	c.JSON(nil, err)
}

// Update 修改文章
func (t *WxArticleServer) Update(c *box.Context) {
	var param model.WxArticle
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}

	param.WxPassportId = c.User.WxPassportId

	err := t.service.WxArticle.Update(&param)
	c.JSON(nil, err)
}

// UpdateCommentSwitch 修改文章评论开关
func (t *WxArticleServer) UpdateCommentSwitch(c *box.Context) {
	var param model.WxArticle
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}

	param.WxPassportId = c.User.WxPassportId

	err := t.service.WxArticle.UpdateCommentSwitch(&param)
	c.JSON(nil, err)
}

// FindByOffiaccount 查询某公众号文章
func (t *WxArticleServer) FindByOffiaccount(c *box.Context) {
	var param model.WxArticle
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}

	r, err := t.service.WxArticle.FindByOffiaccount(&param)
	c.JSON(r, err)
}

// Find 查询文章根据id
func (t *WxArticleServer) Find(c *box.Context) {
	var param model.WxArticle
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}

	r, err := t.service.WxArticle.Find(param.ID)
	c.JSON(r, err)
}

// FindAllByWxPassport 查询某人的全部文章
func (t *WxArticleServer) FindAllByWxPassport(c *box.Context) {
	var param model.WxArticle
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}

	r, err := t.service.WxArticle.FindAllByWxPassport(&param)
	c.JSON(r, err)
}
