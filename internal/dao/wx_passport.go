package dao

import (
	"shippo-server/internal/model"
)

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

	var omits []string

	if m.MiniProgramOpenId == "" {
		omits = append(omits, "mini_program_open_id")
	}

	if m.OffiaccountOpenId == "" {
		omits = append(omits, "offiaccount_open_id")
	}

	if m.Nickname == "" {
		omits = append(omits, "nickname")
	}

	if m.AvatarUrl == "" {
		omits = append(omits, "avatar_url")
	}

	err = t.db.Omit(omits...).Create(&r).Error

	return
}

// FindByUnionId 根据 UnionId 查找微信通行证
func (t *WxPassportDao) FindByUnionId(m *model.WxPassport) (r *model.WxPassport, err error) {
	err = t.db.Where("union_id", m.UnionId).First(&r).Error
	return
}

// Find 查找微信通行证
func (t *WxPassportDao) Find(m *model.WxPassport) (r *model.WxPassport, err error) {
	err = t.db.First(&r, m.ID).Error
	return
}

// Update 修改文章
func (t *WxPassportDao) Update(m *model.WxPassport) (err error) {

	var selects = []string{
		"updated_at",
	}

	if m.MiniProgramOpenId != "" {
		selects = append(selects, "mini_program_open_id")
	}

	if m.OffiaccountOpenId != "" {
		selects = append(selects, "offiaccount_open_id")
	}

	if m.Nickname != "" {
		selects = append(selects, "nickname")
	}

	if m.AvatarUrl != "" {
		selects = append(selects, "avatar_url")
	}

	err = t.db.Select(selects).Updates(&m).Error
	return
}
