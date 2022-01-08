package service

import (
	"shippo-server/internal/dao"
	"time"
)

type Service struct {
	dao                    *dao.Dao
	wxAccessToken          string
	wxAccessTokenCreatedAt time.Time
}

func New() (s *Service) {
	s = &Service{
		dao: dao.New(),
	}

	return s
}
