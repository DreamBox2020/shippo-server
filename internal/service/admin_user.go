package service

import (
	"shippo-server/internal/model"
	"shippo-server/utils/box"
	"shippo-server/utils/check"
	"shippo-server/utils/ecode"
)

func (s *Service) AdminUserCreateEmail(c *box.Context, emial string) (u model.User, err error) {

	if !check.CheckQQEmail(emial) {
		err = ecode.ServerErr
		return
	}

	u, err = s.dao.UserCreateEmail(emial)

	return
}
