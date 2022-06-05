package dao

import (
	"shippo-server/internal/model"
)

type WxArticle struct {
	*Dao
}

func NewWxArticle(s *Dao) *WxArticle {
	return &WxArticle{s}
}

// Create 新增文章
func (t *WxArticle) Create(m *model.WxArticle) (r *model.WxArticle, err error) {
	r = &model.WxArticle{
		Title:         m.Title,
		Url:           m.Url,
		Image1:        m.Image1,
		Image2:        m.Image2,
		OffiaccountId: m.OffiaccountId,
		CommentSwitch: 1,
		WxPassportId:  m.WxPassportId,
	}
	if m.Url == "" {
		err = t.db.Omit("url").Create(r).Error
	} else {
		err = t.db.Create(r).Error
	}
	return
}

// Delete 删除文章
func (t *WxArticle) Delete(id uint) (err error) {
	err = t.db.Delete(&model.WxArticle{}, id).Error
	return
}

// Update 修改文章
func (t *WxArticle) Update(m *model.WxArticle) (err error) {
	tx := t.db.Select("title", "image1", "image2", "offiaccount_id", "updated_at")
	if m.Url != "" {
		tx = tx.Select("url")
	}
	err = tx.Updates(m).Error
	return
}

// UpdateCommentSwitch 修改文章评论开关
func (t *WxArticle) UpdateCommentSwitch(m *model.WxArticle) (err error) {
	err = t.db.Select("comment_switch", "updated_at").Updates(m).Error
	return
}

// FindByOffiaccount 查询某公众号文章
func (t *WxArticle) FindByOffiaccount(m *model.WxArticle) (r *[]model.WxArticle, err error) {
	err = t.db.Model(&model.WxArticle{}).Where("offiaccount_id", m.OffiaccountId).Find(&r).Error
	return
}

// Find 查询文章根据id
func (t *WxArticle) Find(id uint) (r *model.WxArticle, err error) {
	err = t.db.Model(&model.WxArticle{}).First(&r, id).Error
	return
}

// FindAllByWxPassport 查询某人的全部文章
func (t *WxArticle) FindAllByWxPassport(m *model.WxArticle) (r *[]model.WxArticle, err error) {
	err = t.db.Model(&model.WxArticle{}).Where("wx_passport_id", m.WxPassportId).Find(&r).Error
	return
}
