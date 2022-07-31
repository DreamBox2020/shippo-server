package service

import (
	"errors"
	"gorm.io/gorm"
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

func (t *UserService) UserLogin(param model.UserLoginParam, passport model.Passport) (
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

	captcha, err := t.dao.Captcha.FindByTargetAndCode(&model.Captcha{
		Target:  target,
		Code:    param.Code,
		Token:   passport.Token,
		Channel: "login",
	})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = ecode.CaptchaError
		}
		return
	}

	// 如果短信验证码失效，则报错
	if captcha.IsExpire() {
		err = ecode.CaptchaIsExpire
		return
	}

	// 过期验证码
	err = t.dao.Captcha.CaptchaDel(captcha.Target)
	if err != nil {
		return
	}

	// 如果是手机号登陆
	if param.Phone != "" {
		user, err = t.dao.User.UserFindByPhone(captcha.Target)
		if err != nil {
			return
		}

		// 如果没有注册，便自动注册
		if user.Phone == "" {
			user, err = t.dao.User.UserCreate(captcha.Target)
			if err != nil {
				return
			}
		}
	} else {

		user, err = t.dao.User.UserFindByEmail(captcha.Target)
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
	err = t.dao.Passport.PassportUpdate(passport.Token, model.Passport{
		UserId: user.ID,
	})
	if err != nil {
		return
	}

	// 如果当前用户没有绑定微信通行证，且当前通行证中，含有微信通行证，则进行绑定。
	if user.WxPassportId == 0 && passport.WxPassportId != 0 {
		user.WxPassportId = passport.WxPassportId
		err = t.dao.User.UpdateUserWxPassportId(user)
		if err != nil {
			return
		}
	}

	return

}

func (t *UserService) UserFindByUID(uid uint) (u model.User, err error) {
	u, err = t.dao.User.UserFindByUID(uid)
	return
}

func (t *UserService) UserFindByPhone(phone string) (u model.User, err error) {
	u, err = t.dao.User.UserFindByPhone(phone)
	return
}

func (t *UserService) UserFindByEmail(email string) (u model.User, err error) {
	u, err = t.dao.User.UserFindByEmail(email)
	return
}

func (t *UserService) UserFindByWxPassportId(id uint) (u model.User, err error) {
	u, err = t.dao.User.UserFindByWxPassportId(id)
	return
}

func (t *UserService) FindAll(u model.UserFindAllReq) (m model.UserFindAllResp, err error) {
	m, err = t.dao.User.FindAll(u)
	if err != nil {
		return
	}
	for i, v := range m.Items {
		m.Items[i].Phone = utils.PhoneMasking(v.Phone)
		m.Items[i].Email = utils.QQEmailMasking(v.Email)
	}
	return
}

func (t *UserService) UpdateUserRole(u model.User) error {
	return t.dao.User.UpdateUserRole(u)
}

func (t *UserService) UserCreateEmail(email string) (u model.User, err error) {

	if !check.CheckQQEmail(email) {
		err = ecode.ServerErr
		return
	}

	u, err = t.dao.User.UserCreateEmail(email)

	return
}

func (t *UserService) FindByPhone(u *model.User) (r *model.User, err error) {
	r, err = t.dao.User.FindByPhone(u)
	return
}

func (t *UserService) FindByEmail(u *model.User) (r *model.User, err error) {
	r, err = t.dao.User.FindByEmail(u)
	return
}

// PhoneIsRegistered 手机号是否已经注册
func (t *UserService) PhoneIsRegistered(phone string) (bool, error) {
	_, err := t.dao.User.FindByPhone(&model.User{Phone: phone})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// EmailIsRegistered 邮箱是否已经注册
func (t *UserService) EmailIsRegistered(email string) (bool, error) {
	_, err := t.dao.User.FindByEmail(&model.User{Email: email})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
