package dao

import (
	"shippo-server/internal/model"
)

type WxArticleDao struct {
	*Dao
}

func NewWxArticleDao(s *Dao) *WxArticleDao {
	return &WxArticleDao{s}
}

// Create 新增文章
func (t *WxArticleDao) Create(m *model.WxArticle) (r *model.WxArticle, err error) {
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
func (t *WxArticleDao) Delete(id uint) (err error) {
	err = t.db.Delete(&model.WxArticle{}, id).Error
	return
}

// Update 修改文章
func (t *WxArticleDao) Update(m *model.WxArticle) (err error) {
	var selects = []string{
		"title", "image1", "image2", "offiaccount_id", "updated_at",
	}
	if m.Url != "" {
		selects = append(selects, "url")
	}
	err = t.db.Select(selects).Updates(m).Error
	return
}

// UpdateCommentSwitch 修改文章评论开关
func (t *WxArticleDao) UpdateCommentSwitch(m *model.WxArticle) (err error) {
	err = t.db.Select("comment_switch", "updated_at").Updates(m).Error
	return
}

// FindByOffiaccount 查询某公众号文章
func (t *WxArticleDao) FindByOffiaccount(m *model.WxArticle) (r *[]model.WxArticle, err error) {
	err = t.db.Model(&model.WxArticle{}).Where("offiaccount_id", m.OffiaccountId).Find(&r).Error
	return
}

// Find 查询文章根据id
func (t *WxArticleDao) Find(id uint) (r *model.WxArticleExtOffiaccountNickname, err error) {
	subQuery := t.db.Model(&model.WxOffiaccount{})
	err = t.db.Model(&model.WxArticle{}).
		Select("shippo_wx_article.*", "temp.nickname AS offiaccountNickname").
		Joins("Left JOIN (?) temp ON temp.id = offiaccount_id", subQuery).First(&r, id).Error
	return
}

// FindAllByWxPassport 查询某人的全部文章
func (t *WxArticleDao) FindAllByWxPassport(m *model.WxArticle) (
	r *[]model.WxArticleExtOffiaccountNickname, err error) {
	subQuery := t.db.Model(&model.WxOffiaccount{})
	err = t.db.Model(&model.WxArticle{}).
		Select("shippo_wx_article.*", "temp.nickname AS offiaccountNickname").
		Joins("Left JOIN (?) temp ON temp.id = offiaccount_id", subQuery).
		Order("created_at DESC").
		Where("wx_passport_id", m.WxPassportId).Find(&r).Error
	return
}

func (t *WxArticleDao) FindAllByWxPassportAndComment(m *model.WxArticle) (
	r *[]model.WxArticleExtOffiaccountNickname, err error) {
	subQuery := t.db.Model(&model.WxOffiaccount{})
	subQuery2 := t.db.Model(&model.WxComment{}).
		Select("article_id").
		Where("wx_passport_id", m.WxPassportId).
		Where("reply_comment_id IS NULL").
		Group("article_id")

	err = t.db.Model(&model.WxArticle{}).
		Select("shippo_wx_article.*", "temp.nickname AS offiaccountNickname").
		Joins("Left JOIN (?) temp ON temp.id = offiaccount_id", subQuery).
		Order("created_at DESC").
		Where("shippo_wx_article.id IN (?)", subQuery2).Find(&r).Error
	return
}
