package dao

import "shippo-server/internal/model"

type WxPassportDao struct {
	*Dao
}

func NewWxPassportDao(s *Dao) *WxPassportDao {
	return &WxPassportDao{s}
}

// Create 创建微信通行证
func (t *WxPassportDao) Create(m *model.WxPassport) (r *model.WxPassport, err error) {
	r = &model.WxPassport{
		UnionId:           m.UnionId,
		MiniProgramOpenId: m.MiniProgramOpenId,
		OffiaccountOpenId: m.OffiaccountOpenId,
		Nickname:          m.Nickname,
		AvatarUrl:         m.AvatarUrl,
	}

	tx := t.db

	if m.MiniProgramOpenId == "" {
		tx = tx.Omit("mini_program_open_id")
	}

	if m.OffiaccountOpenId == "" {
		tx = tx.Omit("offiaccount_open_id")
	}

	if m.Nickname == "" {
		tx = tx.Omit("nickname")
	}

	if m.AvatarUrl == "" {
		tx = tx.Omit("avatar_url")
	}

	err = tx.Create(&r).Error

	return
}

// FindByUnionId 根据 UnionId 查找微信通行证
func (t *WxPassportDao) FindByUnionId(m *model.WxPassport) (r *model.WxPassport, err error) {
	err = t.db.Where("union_id", m.UnionId).Find(&r).Error
	return
}

// Find 查找微信通行证
func (t *WxPassportDao) Find(m *model.WxPassport) (r *model.WxPassport, err error) {
	err = t.db.Find(&r, m.ID).Error
	return
}

// Update 修改文章
func (t *WxPassportDao) Update(m *model.WxPassport) (err error) {
	tx := t.db.Select("updated_at")

	if m.MiniProgramOpenId != "" {
		tx = tx.Select("mini_program_open_id")
	}

	if m.OffiaccountOpenId != "" {
		tx = tx.Select("offiaccount_open_id")
	}

	if m.Nickname != "" {
		tx = tx.Select("nickname")
	}

	if m.AvatarUrl != "" {
		tx = tx.Select("avatar_url")
	}

	err = tx.Updates(&m).Error
	return
}
