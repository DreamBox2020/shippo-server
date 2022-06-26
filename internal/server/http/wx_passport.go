package http

import (
	"fmt"
	"shippo-server/internal/model"
	"shippo-server/utils/box"
	"shippo-server/utils/ecode"
)

type WxPassportServer struct {
	*Server
	router *box.RouterGroup
}

func NewWxPassportServer(server *Server) *WxPassportServer {
	var s = &WxPassportServer{
		Server: server,
		router: server.router.Group("wxPassport"),
	}
	s.initRouter()
	return s
}

func (t *WxPassportServer) initRouter() {
	t.router.POST("find", t.Find)
	t.router.POST("updateInfo", t.UpdateInfo)
}

// Find 查询自己的微信通行证信息
func (t *WxPassportServer) Find(c *box.Context) {

	if c.User.WxPassportId == 0 {
		c.JSON(nil, ecode.WxPassportIsNull)
		return
	}

	var param = model.WxPassport{}
	param.ID = c.User.WxPassportId

	r, err := t.service.WxPassport.Find(&param)
	c.JSON(r, err)
}

// UpdateInfo 根据code从微信获取用户头像和昵称
func (t *WxPassportServer) UpdateInfo(c *box.Context) {

	if c.User.WxPassportId == 0 {
		c.JSON(nil, ecode.WxPassportIsNull)
		return
	}

	var param struct {
		Code string `json:"code"`
	}
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}
	fmt.Printf("c.ShouldBindJSON->param:%+v\n", param)

	err := t.service.WxPassport.UpdateInfo(param.Code, c.User.WxPassportId)
	c.JSON(nil, err)
}
