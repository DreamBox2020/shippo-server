package service

import (
	"errors"
	"gorm.io/gorm"
	"shippo-server/internal/model"
	"shippo-server/utils/ecode"
)

type WxOffiaccountService struct {
	*Service
}

func NewWxOffiaccountService(s *Service) *WxOffiaccountService {
	return &WxOffiaccountService{s}
}

// FindAll 查询所有公众号
func (t *WxOffiaccountService) FindAll() (r *[]model.WxOffiaccount, err error) {
	r, err = t.dao.WxOffiaccount.FindAll()
	return
}

// FindByUsername 根据username查询公众号
func (t *WxOffiaccountService) FindByUsername(username string) (r *model.WxOffiaccount, err error) {
	r, err = t.dao.WxOffiaccount.FindByUsername(username)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = ecode.ErrRecordNotFound
	}
	return
}

// Find 根据id查询公众号
func (t *WxOffiaccountService) Find(id uint) (r *model.WxOffiaccount, err error) {
	r, err = t.dao.WxOffiaccount.Find(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = ecode.ErrRecordNotFound
	}
	return
}
