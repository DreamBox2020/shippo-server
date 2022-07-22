package service

import (
	"errors"
	"gorm.io/gorm"
	"shippo-server/internal/model"
	"shippo-server/utils/check"
	"shippo-server/utils/ecode"
)

type PassportService struct {
	*Service
}

func NewPassportService(s *Service) *PassportService {
	return &PassportService{s}
}

func (t *PassportService) WxCreate(p model.Passport, code string) (r model.Passport, user model.User, err error) {
	// 获取UnionId
	session, err := t.Group.Wx.AuthCodeToSession(code)
	if err != nil {
		// oauth_code已使用 一般是前端没有处理好导致的，这里后端也做一下兼容。
		if session.Errcode == 40163 {
			return t.PassportCreate(p)
		}
		return
	}

	// 查询绑定该UnionId的微信通行证
	wxPassport, err := t.dao.WxPassport.FindByUnionId(&model.WxPassport{UnionId: session.Unionid})
	if err != nil {
		// 如果没有找到相关通行，就创建一个
		if errors.Is(err, gorm.ErrRecordNotFound) {
			wxPassport, err = t.dao.WxPassport.Create(&model.WxPassport{
				UnionId:           session.Unionid,
				MiniProgramOpenId: session.Openid,
			})
			r, err = t.CreateNoLoginPassport(model.Passport{
				Ip:           p.Ip,
				Ua:           p.Ua,
				Client:       p.Client,
				WxPassportId: wxPassport.ID,
			})
		}
		return
	}

	// 查询绑定该微信通行证的用户
	user, err = t.dao.User.UserFindByWxPassportId(wxPassport.ID)
	if err != nil {
		return
	}

	// 如果没有查到，该通行证可能被解绑过。
	if user.ID == 0 {
		r, err = t.CreateNoLoginPassport(model.Passport{
			Ip:           p.Ip,
			Ua:           p.Ua,
			Client:       p.Client,
			WxPassportId: wxPassport.ID,
		})
		return
	}

	// 创建一个含有登录信息的通行证
	r, err = t.CreateLoginPassport(model.Passport{
		UserId:       user.ID,
		Ip:           p.Ip,
		Ua:           p.Ua,
		Client:       p.Client,
		WxPassportId: wxPassport.ID,
	})

	return
}

func (t *PassportService) PassportCreate(p model.Passport) (r model.Passport, user model.User, err error) {

	// 如果不存在或者失效，就创建一个新的通行证，否则，就续期旧的。
	if p.Token == "" || p.IsExpire() {
		r, err = t.CreateNoLoginPassport(model.Passport{
			Ip:     p.Ip,
			Ua:     p.Ua,
			Client: p.Client,
		})
	} else {
		// 更新ip和ua
		err = t.dao.Passport.PassportUpdate(p.Token, model.Passport{Ip: p.Ip, Ua: p.Ua})
		if err != nil {
			return
		}
		r, err = t.dao.Passport.PassportGet(p.Token)
		if err != nil {
			return
		}
		user, err = t.dao.User.UserFindByUID(r.UserId)
	}
	return
}

func (t *PassportService) PassportGet(passport string) (p model.Passport, err error) {
	if !check.CheckPassport(passport) {
		err = ecode.ServerErr
		return
	}
	return t.dao.Passport.PassportGet(passport)
}

func (t *PassportService) CreateNoLoginPassport(m model.Passport) (r model.Passport, err error) {
	r, err = t.dao.Passport.PassportCreate(model.Passport{
		UserId:       0,
		Ip:           m.Ip,
		Ua:           m.Ua,
		Client:       m.Client,
		WxPassportId: m.WxPassportId,
	})
	return
}

func (t *PassportService) CreateLoginPassport(m model.Passport) (r model.Passport, err error) {
	r, err = t.dao.Passport.PassportCreate(model.Passport{
		UserId:       m.UserId,
		Ip:           m.Ip,
		Ua:           m.Ua,
		Client:       m.Client,
		WxPassportId: m.WxPassportId,
	})
	return
}
