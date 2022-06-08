package http

import (
	"github.com/gin-gonic/gin"
	"shippo-server/internal/model"
	"shippo-server/utils/box"
)

type WxCommentLikeServer struct {
	*Server
	router *gin.RouterGroup
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
	t.router.POST("create", box.Handler(t.Create))
	t.router.POST("delete", box.Handler(t.Delete))
}

// Create 新增点赞
func (t *WxCommentLikeServer) Create(c *box.Context) {
	var m model.WxCommentLike
	c.ShouldBindJSON(&m)

	m.WxPassportId = c.User.WxPassportId
	_, err := t.service.WxCommentLike.Create(&m)

	c.JSON(nil, err)
}

// Delete 取消点赞
func (t *WxCommentLikeServer) Delete(c *box.Context) {
	var m model.WxCommentLike
	c.ShouldBindJSON(&m)

	m.WxPassportId = c.User.WxPassportId
	err := t.service.WxCommentLike.Delete(&m)

	c.JSON(nil, err)

}
