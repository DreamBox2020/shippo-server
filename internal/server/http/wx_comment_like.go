package http

import (
	"github.com/gin-gonic/gin"
	"shippo-server/internal/model"
	"shippo-server/utils/box"
)

type WxCommentLikeServer struct {
	*Server
}

func NewWxCommentLikeServer(s *Server) *WxCommentLikeServer {
	return &WxCommentLikeServer{s}
}

func (t *WxCommentLikeServer) InitRouter(Router *gin.RouterGroup) {
	r := Router.Group("wxCommentLike")
	{
		r.POST("create", box.Handler(t.Create))
		r.POST("delete", box.Handler(t.Delete))
	}
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
