package service

import (
	"shippo-server/internal/model"
	"shippo-server/utils/check"
	"shippo-server/utils/ecode"
	"time"
)

type PassportService struct {
	*Service
}

func NewPassportService(s *Service) *PassportService {
	return &PassportService{s}
}

func (s *PassportService) PassportCreate(p model.Passport) (data model.PassportCreateResult, err error) {

	// 如果不存在或者到期(30天)，就创建一个新的通行证，否则，就续期旧的。
	if p.Token == "" || time.Since(p.UpdatedAt) > time.Hour*24*30 {

		p, err = s.dao.Passport.PassportCreate(model.Passport{
			Token:  "",
			UserId: 0,
			Ip:     p.Ip,
			Ua:     p.Ua,
			Client: p.Client,
		})
		if err != nil {
			return
		}

	} else {

		// 更新ip和ua
		p, err = s.dao.Passport.PassportUpdate(p.Token, model.Passport{Ip: p.Ip, Ua: p.Ua})
		if err != nil {
			return
		}
	}

	data.Passport = p.Token
	data.Uid = p.UserId

	var access []model.PermissionAccess
	var user model.User

	if p.UserId == 0 {
		// 如果当前没有登录，查询基础权限信息
		access, err = s.Group.PermissionPolicy.FindPermissionAccessByPolicyName("SysBase")
		if err != nil {
			return
		}
	} else {
		// 如果当前登录，就获取用户信息
		user, err = s.Group.User.UserFindByUID(p.UserId)
		if err != nil {
			return
		}

		// 根据用户角色查询对应权限信息
		access, err = s.Group.Role.RoleFindPermissionAccess(user.Role)
		if err != nil {
			return
		}
	}

	data.Access = access

	return
}

func (s *PassportService) PassportGet(passport string) (p model.Passport, err error) {
	if !check.CheckPassport(passport) {
		err = ecode.ServerErr
		return
	}
	return s.dao.Passport.PassportGet(passport)
}
