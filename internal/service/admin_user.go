package service

import (
	"shippo-server/internal/model"
	"shippo-server/utils/check"
	"shippo-server/utils/ecode"
)

type AdminUserService struct {
	*Service
}

func NewAdminUserService(s *Service) *AdminUserService {
	return &AdminUserService{s}
}

func (s *AdminUserService) AdminUserCreateEmail(emial string) (u model.User, err error) {

	if !check.CheckQQEmail(emial) {
		err = ecode.ServerErr
		return
	}

	u, err = s.dao.User.UserCreateEmail(emial)

	return
}
