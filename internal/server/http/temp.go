package http

import (
	"fmt"
	"shippo-server/internal/model"
	"shippo-server/utils/box"
)

type TempServer struct {
	*Server
	router *box.RouterGroup
}

func NewTempServer(server *Server) *TempServer {
	var s = &TempServer{
		Server: server,
		router: server.router.Group("temp"),
	}
	s.initRouter()
	return s
}

func (t *TempServer) initRouter() {
	t.router.POST("temp_trade_20220108/find", t.Temp_trade_20220108_find)
	t.router.POST("temp_trade_20220108/add", t.Temp_trade_20220108_add)
	t.router.POST("temp_trade_20220108/findNoExist", t.Temp_trade_20220108_findNoExist)
}

func (t *TempServer) Temp_trade_20220108_find(c *box.Context) {

	var param = new(struct {
		Qq string `json:"qq"`
		Id string `json:"id"`
	})
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}
	fmt.Printf("c.ShouldBindJSON->param:%+v\n", param)

	// 如果参数中含有QQ，那么就按照QQ查找，否则按照订单号。
	if param.Qq != "" {
		data, err := t.service.Temp.Temp_trade_20220108_findByUserQQ(param.Qq)
		c.JSON(data, err)
	} else {
		data, err := t.service.Temp.Temp_trade_20220108_findByTradeId(param.Id)
		c.JSON(data, err)
	}
}

func (t *TempServer) Temp_trade_20220108_add(c *box.Context) {
	var param model.Temp_trade_20220108_TradeAddParam
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}
	fmt.Printf("c.ShouldBindJSON->param:%+v\n", param)

	data, err := t.service.Temp.Temp_trade_20220108_add(param)
	c.JSON(data, err)
}

func (t *TempServer) Temp_trade_20220108_findNoExist(c *box.Context) {
	var param = new(struct {
		List []string `json:"list"`
	})
	if err := c.ShouldBindJSON(&param); err != nil {
		return
	}
	fmt.Printf("c.ShouldBindJSON->param:%+v\n", param)

	data, err := t.service.Temp.Temp_trade_20220108_findNoExist(param.List)
	c.JSON(data, err)
}
