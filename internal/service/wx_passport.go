package service

import (
	"errors"
	"gorm.io/gorm"
	"shippo-server/internal/model"
	"shippo-server/utils/ecode"
)

type WxPassportService struct {
	*Service
}

func NewWxPassportService(s *Service) *WxPassportService {
	return &WxPassportService{s}
}

// Create 创建微信通行证
func (t *WxPassportService) Create(m *model.WxPassport) (r *model.WxPassport, err error) {
	r, err = t.dao.WxPassport.Create(m)
	return
}

// FindByUnionId 根据 UnionId 查找微信通行证
func (t *WxPassportService) FindByUnionId(m *model.WxPassport) (r *model.WxPassport, err error) {
	r, err = t.dao.WxPassport.FindByUnionId(m)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = ecode.ErrRecordNotFound
	}
	return
}

// Find 查找微信通行证
func (t *WxPassportService) Find(m *model.WxPassport) (r *model.WxPassport, err error) {
	r, err = t.dao.WxPassport.Find(m)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = ecode.ErrRecordNotFound
	}
	return
}

// Update 修改文章
func (t *WxPassportService) Update(m *model.WxPassport) (err error) {
	err = t.dao.WxPassport.Update(m)

	return
}
