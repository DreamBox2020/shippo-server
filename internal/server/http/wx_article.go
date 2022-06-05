package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shippo-server/internal/model"
	"shippo-server/utils/box"
	"shippo-server/utils/ecode"
)

type WxArticle struct {
	*Server
}

func NewWxArticle(s *Server) *WxArticle {
	return &WxArticle{s}
}
func (t *WxArticle) InitRouter(Router *gin.RouterGroup) {
	r := Router.Group("wxArticle")
	{
		r.POST("create", box.Handler(t.Create))
		r.POST("update", box.Handler(t.Update))
		r.POST("updateCommentSwitch", box.Handler(t.UpdateCommentSwitch))
		r.POST("findByOffiaccount", box.Handler(t.FindByOffiaccount))
		r.POST("find", box.Handler(t.Find))
		r.POST("findAllByWxPassport", box.Handler(t.FindAllByWxPassport))
	}
}

// Create 新增文章
func (t *WxArticle) Create(c *box.Context) {
	if c.User.WxPassportId == 0 {
		c.JSON(nil, ecode.WxPassportIsNull)
		return
	}

	var m model.WxArticle
	c.ShouldBindJSON(&m)

	m.WxPassportId = c.User.WxPassportId

	fmt.Printf("WxArticle->Create:%+v", m)

	_, err := t.service.WxArticle.Create(&m)
	c.JSON(nil, err)
}

// Update 修改文章
func (t *WxArticle) Update(c *box.Context) {
	var m model.WxArticle
	c.ShouldBindJSON(&m)

	m.WxPassportId = c.User.WxPassportId

	fmt.Printf("WxArticle->Update:%+v", m)

	err := t.service.WxArticle.Update(&m)
	c.JSON(nil, err)
}

// UpdateCommentSwitch 修改文章评论开关
func (t *WxArticle) UpdateCommentSwitch(c *box.Context) {
	var m model.WxArticle
	c.ShouldBindJSON(&m)

	m.WxPassportId = c.User.WxPassportId
	fmt.Printf("WxArticle->UpdateCommentSwitch:%+v", m)

	err := t.service.WxArticle.UpdateCommentSwitch(&m)
	c.JSON(nil, err)
}

// FindByOffiaccount 查询某公众号文章
func (t *WxArticle) FindByOffiaccount(c *box.Context) {
	var m model.WxArticle
	c.ShouldBindJSON(&m)

	fmt.Printf("WxArticle->FindByOffiaccount:%+v\n", m)

	r, err := t.service.WxArticle.FindByOffiaccount(&m)
	c.JSON(r, err)
}

// Find 查询文章根据id
func (t *WxArticle) Find(c *box.Context) {
	var m model.WxArticle
	c.ShouldBindJSON(&m)

	fmt.Printf("WxArticle->Find:%+v\n", m)

	r, err := t.service.WxArticle.Find(m.ID)
	c.JSON(r, err)
}

// FindAllByWxPassport 查询某人的全部文章
func (t *WxArticle) FindAllByWxPassport(c *box.Context) {
	var m model.WxArticle
	c.ShouldBindJSON(&m)

	fmt.Printf("WxArticle->FindAllByWxPassport:%+v\n", m)

	r, err := t.service.WxArticle.FindAllByWxPassport(&m)
	c.JSON(r, err)
}
