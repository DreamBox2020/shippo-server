package service

import "shippo-server/internal/dao"

type Service struct {
	dao *dao.Dao
}

func New() (s *Service) {
	s = &Service{
		dao: dao.New(),
	}

	return s
}
