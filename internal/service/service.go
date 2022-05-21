package service

import (
	"shippo-server/internal/dao"
	"time"
)

type ServiceGroup struct {
	User             *UserService
	Wx               *WxService
	Temp             *TempService
	Passport         *PassportService
	Captcha          *CaptchaService
	AdminUser        *AdminUserService
	Role             *RoleService
	PermissionAccess *PermissionAccessService
	PermissionPolicy *PermissionPolicyService
	Picture          *PictureService
}

type Service struct {
	dao                    *dao.DaoGroup
	Group                  *ServiceGroup
	wxAccessToken          string
	wxAccessTokenCreatedAt time.Time
}

func New() *Service {
	var d = dao.New()
	s := &Service{
		dao:   d.Group,
		Group: nil,
	}
	s.Group = NewGroup(s)

	return s
}

func NewGroup(d *Service) *ServiceGroup {
	return &ServiceGroup{
		User:             NewUserService(d),
		Wx:               NewWxService(d),
		Temp:             NewTempService(d),
		Passport:         NewPassportService(d),
		Captcha:          NewCaptchaService(d),
		AdminUser:        NewAdminUserService(d),
		Role:             NewRoleService(d),
		PermissionAccess: NewPermissionAccessService(d),
		PermissionPolicy: NewPermissionPolicyService(d),
		Picture:          NewPictureService(d),
	}
}
