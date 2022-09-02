package service

import (
	"shippo-server/utils/check"
	"shippo-server/utils/ecode"
	email2 "shippo-server/utils/email"
	"shippo-server/utils/sms"
)

type CaptchaService struct {
	*Service
}

func NewCaptchaService(s *Service) *CaptchaService {
	return &CaptchaService{s}
}

func (t *CaptchaService) CaptchaSmsSend(phone string, token string, channel string) (err error) {

	if !check.CheckPhone(phone) {
		err = ecode.ServerErr
		return
	}

	// bind 绑定 rebind-set 更换绑定，设置新绑定
	// 必须没有被注册
	if channel == "bind" || channel == "rebind-set" {
		var tag bool
		tag, err = t.Group.User.PhoneIsRegistered(phone)
		if err != nil {
			return
		}

		if tag {
			err = ecode.AccountRegistered
			return
		}
	}

	// rebind-verify 更换绑定，验证旧绑定
	// 必须已经注册
	if channel == "rebind-verify" {
		var tag bool
		tag, err = t.Group.User.PhoneIsRegistered(phone)
		if err != nil {
			return
		}

		if !tag {
			err = ecode.AccountUnregistered
			return
		}
	}

	// 过期所有验证码
	err = t.dao.Captcha.CaptchaDel(phone)
	if err != nil {
		return
	}

	// 生成新的验证码
	r, err := t.dao.Captcha.CaptchaSmsInsert(phone, token, channel)
	if err != nil {
		return
	}

	// 发送验证码
	if !sms.SendSms(r.Target, r.Code) {
		return ecode.CaptchaSendError
	}

	return

}

func (t *CaptchaService) CaptchaEmailSend(email string, token string, channel string) (err error) {

	if !check.CheckQQEmail(email) {
		err = ecode.ServerErr
		return
	}

	// bind 绑定 rebind-set 更换绑定，设置新绑定
	// 必须没有被注册
	if channel == "bind" || channel == "rebind-set" {
		var tag bool
		tag, err = t.Group.User.EmailIsRegistered(email)
		if err != nil {
			return
		}

		if tag {
			err = ecode.AccountRegistered
			return
		}
	}

	// rebind-verify 更换绑定，验证旧绑定
	// 必须已经注册
	if channel == "rebind-verify" {
		var tag bool
		tag, err = t.Group.User.EmailIsRegistered(email)
		if err != nil {
			return
		}

		if !tag {
			err = ecode.AccountUnregistered
			return
		}
	}

	// 过期所有验证码
	err = t.dao.Captcha.CaptchaDel(email)
	if err != nil {
		return
	}

	// 生成新的验证码
	r, err := t.dao.Captcha.CaptchaEmailInsert(email, token, channel)
	if err != nil {
		return
	}

	// 发送验证码
	if !email2.SendEmail(r.Target, r.Code) {
		return ecode.CaptchaSendError
	}

	return
}
