package http

import (
	"shippo-server/internal/model"
	"shippo-server/utils/box"
)

type WxCommentLikeServer struct {
	*Server
	router *box.RouterGroup
}

func NewWxCommentLikeServer(server *Server) *WxCommentLikeServer {
	var s = &WxCommentLikeServer{
		Server: server,
		router: server.router.Group("wxCommentLike"),
	}
	s.initRouter()
	return s
}

func (t *WxCommentLikeServer) initRouter() {
	t.router.POST("create", t.Create)
	t.router.POST("delete", t.Delete)
}

// Create 新增点赞
func (t *WxCommentLikeServer) Create(c *box.Context) {
	var param model.WxCommentLike
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}

	param.WxPassportId = c.User.WxPassportId
	_, err := t.service.WxCommentLike.Create(&param)

	c.JSON(nil, err)
}

// Delete 取消点赞
func (t *WxCommentLikeServer) Delete(c *box.Context) {
	var param model.WxCommentLike
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}

	param.WxPassportId = c.User.WxPassportId
	err := t.service.WxCommentLike.Delete(&param)

	c.JSON(nil, err)

}
