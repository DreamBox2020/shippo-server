package service

import (
	"errors"
	"gorm.io/gorm"
	"shippo-server/internal/dao"
	"time"
)

type ServiceGroup struct {
	User             *UserService
	Wx               *WxService
	WxArticle        *WxArticleService
	Temp             *TempService
	Passport         *PassportService
	Captcha          *CaptchaService
	Role             *RoleService
	PermissionAccess *PermissionAccessService
	PermissionPolicy *PermissionPolicyService
	Picture          *PictureService
	WxCommentLike    *WxCommentLikeService
	WxOffiaccount    *WxOffiaccountService
	WxComment        *WxCommentService
	WxPassport       *WxPassportService
	File             *FileService
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
	s.initGroup()

	return s
}

func (t *Service) initGroup() {
	t.Group = &ServiceGroup{
		User:             NewUserService(t),
		Wx:               NewWxService(t),
		WxArticle:        NewWxArticleService(t),
		Temp:             NewTempService(t),
		Passport:         NewPassportService(t),
		Captcha:          NewCaptchaService(t),
		Role:             NewRoleService(t),
		PermissionAccess: NewPermissionAccessService(t),
		PermissionPolicy: NewPermissionPolicyService(t),
		Picture:          NewPictureService(t),
		WxCommentLike:    NewWxCommentLikeService(t),
		WxOffiaccount:    NewWxOffiaccountService(t),
		WxComment:        NewWxCommentService(t),
		WxPassport:       NewWxPassportService(t),
		File:             NewFileService(t),
	}
}

func (t *Service) IsErrRecordNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
