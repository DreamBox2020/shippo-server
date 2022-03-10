package service

import (
	"fmt"
	"shippo-server/internal/model"
	"shippo-server/utils"
	"shippo-server/utils/ecode"
)

type TempService struct {
	*Service
}

func NewTempService(s *Service) *TempService {
	return &TempService{s}
}

func (s *TempService) Temp_trade_20220108_findByTradeId(id string) (data map[string]interface{}, err error) {

	// TODO 检查id是否正确

	// 查询
	r, err := s.dao.Temp.Temp_trade_20220108_findByTradeId(id)

	if err != nil {
		return
	}

	data = make(map[string]interface{})
	data["id"] = r.TradeId
	data["type"] = r.TradeType
	data["amount"] = r.TradeAmount
	data["status"] = r.AmountStatus
	data["qq"] = utils.QQMasking(r.UserQq)
	data["phone"] = utils.PhoneMasking(r.UserPhone)
	data["time"] = utils.FormatTime(r.CreatedAt)

	return
}

func (s *TempService) Temp_trade_20220108_findByUserQQ(qq string) (data []map[string]interface{}, err error) {

	// TODO 检查qq是否正确

	// 查询
	r, err := s.dao.Temp.Temp_trade_20220108_findByUserQQ(qq)

	if err != nil {
		return
	}

	data = make([]map[string]interface{}, len(r))

	for key, value := range r {
		data[key] = make(map[string]interface{})
		//data[key]["id"] = value.TradeId
		data[key]["type"] = value.TradeType
		data[key]["amount"] = value.TradeAmount
		data[key]["status"] = value.AmountStatus
		data[key]["qq"] = utils.QQMasking(value.UserQq)
		data[key]["phone"] = utils.PhoneMasking(value.UserPhone)
		data[key]["time"] = utils.FormatTime(value.CreatedAt)
	}

	return
}

func (s *TempService) Temp_trade_20220108_add(m model.Temp_trade_20220108_TradeAddParam) (data interface{}, err error) {

	// TODO 校验所有参数

	if m.TradeId1 == "" {
		err = ecode.ServerErr
		return
	}

	// 1. 获取 定金的订单信息
	t1, err := s.dao.Temp.Temp_trade_20220108_findByTradeId(m.TradeId1)

	if err != nil {
		return
	}

	fmt.Printf("Temp_trade_20220108_add->t1: %+v\n", t1)

	// 2. 如果没有查到 定金的订单信息，则报错
	if t1.TradeId == "" {
		err = ecode.Temp_trade_20220108_Trade1NotFind
		return
	}

	// 3. 如果 已经绑定所有者 且 qq不一致，则报错
	if t1.UserQq != "" && t1.UserQq != m.UserQq {
		err = ecode.Temp_trade_20220108_Trade1Repeat
		return
	}

	// 3. 如果 订单状态异常，则报错
	if t1.AmountStatus != 0 {
		err = ecode.ServerErr
		return
	}

	// 4. 如果定金 金额不是100，也不是233元，则报错 （防止补款填写到定金）
	if t1.TradeAmount != 100 && t1.TradeAmount != 233 {
		err = ecode.Temp_trade_20220108_Trade1AmountErr
		return
	}

	// 5. 绑定所有者
	t1.UserQq = m.UserQq
	t1.UserPhone = m.UserPhone
	t1, err = s.dao.Temp.Temp_trade_20220108_save(t1)

	if err != nil {
		return
	}

	// 6. 如果金额大于等于233，则不处理第二个订单
	if t1.TradeAmount >= 233 {
		return
	}

	if m.TradeId2 == "" {
		err = ecode.ServerErr
		return
	}

	t2, err := s.dao.Temp.Temp_trade_20220108_findByTradeId(m.TradeId2)

	if err != nil {
		return
	}

	fmt.Printf("Temp_trade_20220108_add->t2: %+v\n", t2)

	if t2.TradeId == "" {
		err = ecode.Temp_trade_20220108_Trade2NotFind
		return
	}

	if t2.UserQq != "" && t2.UserQq != m.UserQq {
		err = ecode.Temp_trade_20220108_Trade2Repeat
		return
	}

	if t2.AmountStatus != 0 {
		err = ecode.ServerErr
		return
	}

	if t2.TradeAmount != 133 {
		err = ecode.Temp_trade_20220108_Trade2AmountErr
		return
	}

	t2.UserQq = m.UserQq
	t2.UserPhone = m.UserPhone
	t2, err = s.dao.Temp.Temp_trade_20220108_save(t2)

	if err != nil {
		return
	}

	return
}

func (s *TempService) Temp_trade_20220108_findNoExist(list []string) (data []string, err error) {
	// 1. 查询出订单金额 >= 233；订单状态为（0正常）的订单
	t1, err := s.dao.Temp.Temp_trade_20220108_findSuccess()

	if err != nil {
		return
	}
	//遍历查询两个数组,若前台数据未匹配到数据库,则新增到数组QQNoExist
	for i := 0; i < len(list); i++ {
		flag := true
		for _, trade := range t1 {
			if trade.UserQq == list[i] {
				flag = false
				break
			}
		}
		if flag {
			data = append(data, list[i])
		}
	}
	return
}
