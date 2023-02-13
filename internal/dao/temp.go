package dao

import (
	"shippo-server/internal/model"
)

type TempDao struct {
	*Dao
}

func NewTempDao(s *Dao) *TempDao {
	return &TempDao{s}
}

// Temp_trade_20220108_findByTradeId 根据订单号查询订单
func (t *TempDao) Temp_trade_20220108_findByTradeId(id string) (p model.Temp_trade_20220108, err error) {
	err = t.db.Where("trade_id", id).Limit(1).Find(&p).Error
	return
}

// Temp_trade_20220108_findByUserQQ 根据用户QQ查询订单
func (t *TempDao) Temp_trade_20220108_findByUserQQ(qq string) (p []model.Temp_trade_20220108, err error) {
	err = t.db.Where("user_qq", qq).Find(&p).Error
	return
}

// Temp_trade_20220108_save 创建订单
func (t *TempDao) Temp_trade_20220108_save(p model.Temp_trade_20220108) (model.Temp_trade_20220108, error) {
	return p, t.db.Save(&p).Error
}

// Temp_trade_20220108_findSuccess 查询出订单金额 >= 233；订单状态为（0正常）的订单
func (t *TempDao) Temp_trade_20220108_findSuccess() (p []model.Temp_trade_20220108_FindSuccessResult, err error) {
	err = t.db.Model(&model.Temp_trade_20220108{}).Select("user_qq", "sum(trade_amount) as amount").
		Group("user_qq").Where("amount_status", 0).Having("amount>=233").Find(&p).Error
	return
}

// Temp_express_20220914_findByQQAndPhone 根据QQ和手机号查询快递订单号
func (t *TempDao) Temp_express_20220914_findByQQAndPhone(m *model.Temp_express_20220914) (
	r *[]model.Temp_express_20220914, err error) {
	err = t.db.Where("qq", m.QQ).Where("phone", m.Phone).Find(&r).Error
	return
}
