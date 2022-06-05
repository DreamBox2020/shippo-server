package service

import (
	"errors"
	"gorm.io/gorm"
	"shippo-server/internal/model"
	"shippo-server/utils/ecode"
)

type WxArticle struct {
	*Service
}

func NewWxArticle(s *Service) *WxArticle {
	return &WxArticle{s}
}

// Create 新增文章
func (t *WxArticle) Create(m *model.WxArticle) (r *model.WxArticle, err error) {
	r, err = t.dao.WxArticle.Create(m)
	return
}

// Update 修改文章
func (t *WxArticle) Update(m *model.WxArticle) (err error) {
	old, err := t.dao.WxArticle.Find(m.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = ecode.ServerErr
		}
		return
	}

	if old.WxPassportId != m.WxPassportId {
		return ecode.ServerErr
	}

	if old.Url != "" {
		return ecode.WxArticleUpdateProhibit
	}

	err = t.dao.WxArticle.Update(m)
	return
}

// UpdateCommentSwitch 修改文章评论开关
func (t *WxArticle) UpdateCommentSwitch(m *model.WxArticle) (err error) {
	old, err := t.dao.WxArticle.Find(m.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = ecode.ServerErr
		}
		return
	}

	if old.WxPassportId != m.WxPassportId {
		return ecode.ServerErr
	}

	err = t.dao.WxArticle.UpdateCommentSwitch(m)
	return
}

// FindByOffiaccount 查询某公众号文章
func (t *WxArticle) FindByOffiaccount(m *model.WxArticle) (r *[]model.WxArticle, err error) {
	r, err = t.dao.WxArticle.FindByOffiaccount(m)
	return
}

// Find 查询文章根据id
func (t *WxArticle) Find(id uint) (r *model.WxArticle, err error) {
	r, err = t.dao.WxArticle.Find(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = ecode.ErrRecordNotFound
	}
	return
}

// FindAllByWxPassport 查询某人的全部文章
func (t *WxArticle) FindAllByWxPassport(m *model.WxArticle) (r *[]model.WxArticle, err error) {
	r, err = t.dao.WxArticle.FindAllByWxPassport(m)
	return
}
