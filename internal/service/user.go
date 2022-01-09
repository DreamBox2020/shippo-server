package service

import (
	"shippo-server/internal/model"
	"shippo-server/utils/box"
	"shippo-server/utils/check"
	"shippo-server/utils/ecode"
)

func (s *Service) UserLogin(c *box.Context, param model.UserLoginParam, token string) (data map[string]interface{}, err error) {
	var user model.User
	var p model.Passport

	if !check.CheckPassport(token) {
		err = ecode.ServerErr
		return
	}

	if !check.CheckPhone(param.Phone) {
		err = ecode.ServerErr
		return
	}

	if !check.CheckSmsCode(param.Code) {
		err = ecode.ServerErr
		return
	}

	sms, err := s.dao.SmsByPhoneAndCode(param.Phone, param.Code, token)
	if err != nil {
		return
	}

	// 如果短信验证成功
	if sms.Phone != "" {
		// 过期验证码
		s.dao.SmsDel(sms.Phone)

		user, err = s.dao.UserFindByPhone(sms.Phone)
		if err != nil {
			return
		}

		// 如果没有注册
		if user.Phone == "" {
			user, err = s.dao.UserCreate(sms.Phone)
			if err != nil {
				return
			}
		}

		// 更新用户信息
		p, err = s.dao.PassportUpdate(token, model.Passport{
			UserId: user.ID,
		})
		if err != nil {
			return
		}

		data = make(map[string]interface{})
		data["passport"] = p.Token
		data["uid"] = p.UserId

		return
	}

	return
}
