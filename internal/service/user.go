package service

import (
	"shippo-server/internal/model"
	"shippo-server/utils"
	"shippo-server/utils/check"
	"shippo-server/utils/ecode"
)

type UserService struct {
	*Service
}

func NewUserService(s *Service) *UserService {
	return &UserService{s}
}

func (s *UserService) UserLogin(param model.UserLoginParam, passport model.Passport) (
	user model.User, err error) {

	if !check.CheckPassport(passport.Token) {
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

	captcha, err := s.dao.Captcha.CaptchaByTargetAndCode(target, param.Code, passport.Token)
	if err != nil {
		return
	}

	// 如果短信验证失败
	if captcha.Target == "" {
		err = ecode.CaptchaError
		return
	}

	// 过期验证码
	err = s.dao.Captcha.CaptchaDel(captcha.Target)
	if err != nil {
		return
	}

	// 如果是手机号登陆
	if param.Phone != "" {
		user, err = s.dao.User.UserFindByPhone(captcha.Target)
		if err != nil {
			return
		}

		// 如果没有注册，便自动注册
		if user.Phone == "" {
			user, err = s.dao.User.UserCreate(captcha.Target)
			if err != nil {
				return
			}
		}
	} else {

		user, err = s.dao.User.UserFindByEmail(captcha.Target)
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
	err = s.dao.Passport.PassportUpdate(passport.Token, model.Passport{
		UserId: user.ID,
	})
	if err != nil {
		return
	}

	// 如果当前用户没有绑定微信通行证，且当前通行证中，含有微信通行证，则进行绑定。
	if user.WxPassportId == 0 && passport.WxPassportId != 0 {
		user.WxPassportId = passport.WxPassportId
		err = s.dao.User.UpdateUserWxPassportId(user)
		if err != nil {
			return
		}
	}

	return

}

func (s *UserService) UserFindByUID(uid uint) (u model.User, err error) {
	u, err = s.dao.User.UserFindByUID(uid)
	return
}

func (s *UserService) UserFindByPhone(phone string) (u model.User, err error) {
	u, err = s.dao.User.UserFindByPhone(phone)
	return
}

func (s *UserService) UserFindByEmail(email string) (u model.User, err error) {
	u, err = s.dao.User.UserFindByEmail(email)
	return
}

func (s *UserService) UserFindByWxPassportId(id uint) (u model.User, err error) {
	u, err = s.dao.User.UserFindByWxPassportId(id)
	return
}

func (s *UserService) FindAll(u model.UserFindAllReq) (m model.UserFindAllResp, err error) {
	m, err = s.dao.User.FindAll(u)
	if err != nil {
		return
	}
	for i, v := range m.Items {
		m.Items[i].Phone = utils.PhoneMasking(v.Phone)
		m.Items[i].Email = utils.QQEmailMasking(v.Email)
	}
	return
}

func (s *UserService) UpdateUserRole(u model.User) error {
	return s.dao.User.UpdateUserRole(u)
}

func (s *UserService) UserCreateEmail(email string) (u model.User, err error) {

	if !check.CheckQQEmail(email) {
		err = ecode.ServerErr
		return
	}

	u, err = s.dao.User.UserCreateEmail(email)

	return
}
