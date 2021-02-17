package service

import (
	"shippo-server/utils"
	"shippo-server/utils/box"
)

func (s *Service) SmsSend(c *box.Context, phone string, token string) (err error) {

	// 过期所有验证码
	err = s.dao.SmsDel(phone)
	if err != nil {
		return
	}
	// 生成新的验证码
	r, err := s.dao.SmsInsert(phone, token)
	if err != nil {
		return
	}

	// 发送验证码
	utils.SendSms(r.Phone, r.Code)
	return
}
