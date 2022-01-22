package service

import (
	"shippo-server/utils"
	"shippo-server/utils/box"
	"shippo-server/utils/check"
	"shippo-server/utils/ecode"
)

func (s *Service) CaptchaSmsSend(c *box.Context, phone string, token string) (err error) {

	if !check.CheckPhone(phone) {
		err = ecode.ServerErr
		return
	}

	// 过期所有验证码
	err = s.dao.CaptchaDel(phone)
	if err != nil {
		return
	}
	// 生成新的验证码
	r, err := s.dao.CaptchaSmsInsert(phone, token)
	if err != nil {
		return
	}

	// 发送验证码
	utils.SendSms(r.Target, r.Code)
	return
}

func (s *Service) CaptchaEmailSend(c *box.Context, email string, token string) (err error) {

	if !check.CheckQQEmail(email) {
		err = ecode.ServerErr
		return
	}

	// 过期所有验证码
	err = s.dao.CaptchaDel(email)
	if err != nil {
		return
	}
	// 生成新的验证码
	r, err := s.dao.CaptchaEmailInsert(email, token)
	if err != nil {
		return
	}

	// 发送验证码
	utils.SendEmail(r.Target, r.Code)
	return
}
