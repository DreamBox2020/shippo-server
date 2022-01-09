package model

import "gorm.io/gorm"

type Temp_trade_20220108 struct {
	gorm.Model
	TradeId      string
	TradeType    uint
	TradeAmount  uint
	AmountStatus uint
	UserQq       string
	UserPhone    string
}

// 新增订单信息的参数模型
type Temp_trade_20220108_TradeAddParam struct {
	TradeId1  string `json:"trade1"`
	TradeId2  string `json:"trade2"`
	UserQq    string `json:"qq"`
	UserPhone string `json:"phone"`
}
