package dao

import (
	"shippo-server/internal/model"
)

// 根据订单号查询订单
func (d *Dao) Temp_trade_20220108_findByTradeId(id string) (p model.Temp_trade_20220108, err error) {
	err = d.db.Where("trade_id", id).Limit(1).Find(&p).Error
	return
}

// 根据用户QQ查询订单
func (d *Dao) Temp_trade_20220108_findByUserQQ(qq string) (p []model.Temp_trade_20220108, err error) {
	err = d.db.Where("user_qq", qq).Find(&p).Error
	return
}

// 创建订单
func (d *Dao) Temp_trade_20220108_save(p model.Temp_trade_20220108) (model.Temp_trade_20220108, error) {
	return p, d.db.Save(&p).Error
}

// 查询出订单金额 >= 233；订单状态为（0正常）的订单
func (d *Dao) Temp_trade_20220108_findSuccess() (p []model.Temp_trade_20220108_FindSuccessResult, err error) {
	err = d.db.Table("(?) as u",
		d.db.Model(&model.Temp_trade_20220108{}).
			Select("User_qq as UserQq ,sum(trade_amount) as tradeAmount, sum(amount_status) as amountStatus").
			Group("user_qq").Find(&model.Temp_trade_20220108_FindSuccessResult{})).
		Where("tradeAmount>=233 and amountStatus=0").Find(&p).Error
	return
}
