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

	if param.Phone == "" && param.Email == "" {
		err = ecode.ServerErr
		return
	}

	if param.Phone != "" && !check.CheckPhone(param.Phone) {
		err = ecode.ServerErr
		return
	}

	if param.Email != "" && !check.CheckQQEmail(param.Email) {
		err = ecode.ServerErr
		return
	}

	if !check.CheckSmsCode(param.Code) {
		err = ecode.ServerErr
		return
	}

	var target string
	if param.Phone == "" {
		target = param.Email
	} else {
		target = param.Phone
	}

	captcha, err := s.dao.CaptchaByTargetAndCode(target, param.Code, token)
	if err != nil {
		return
	}

	// 如果短信验证成功
	if captcha.Target != "" {
		// 过期验证码
		err = s.dao.CaptchaDel(captcha.Target)
		if err != nil {
			return
		}

		// 如果是手机号登陆
		if param.Phone != "" {
			user, err = s.dao.UserFindByPhone(captcha.Target)
			if err != nil {
				return
			}

			// 如果没有注册
			if user.Phone == "" {
				user, err = s.dao.UserCreate(captcha.Target)
				if err != nil {
					return
				}
			}
		} else {

			user, err = s.dao.UserFindByEmail(captcha.Target)
			if err != nil {
				return
			}

			// 如果邮箱登陆，没有注册，就报错
			if user.Email == "" {
				err = ecode.ServerErr
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
