package dao

import (
	"shippo-server/internal/model"
)

type WxCommentDao struct {
	*Dao
}

func NewWxCommentDao(s *Dao) *WxCommentDao {
	return &WxCommentDao{s}
}

// FindCommentByArticle 查询某文章的全部一级评论
func (t *WxCommentDao) FindCommentByArticle(m *model.WxComment) (r *[]model.WxCommentExt, err error) {
	subQuery := t.db.Model(&model.WxPassport{})
	err = t.db.Model(&model.WxComment{}).
		Select("shippo_wx_comment.*", "temp.nickname AS nickname", "temp.avatar_url AS avatar_url").
		Joins("Left JOIN (?) temp ON temp.id = wx_passport_id", subQuery).
		Order("created_at DESC").
		Where("article_id", m.ArticleId).
		Where("reply_comment_id IS NULL").Find(&r).Error
	return
}

// FindReplyByArticle 查询某文章的全部二级评论
func (t *WxCommentDao) FindReplyByArticle(m *model.WxComment) (r *[]model.WxCommentExt, err error) {
	subQuery := t.db.Model(&model.WxPassport{})
	err = t.db.Model(&model.WxComment{}).
		Select("shippo_wx_comment.*", "temp.nickname AS nickname", "temp.avatar_url AS avatar_url").
		Joins("Left JOIN (?) temp ON temp.id = wx_passport_id", subQuery).
		Order("created_at ASC").
		Where("article_id", m.ArticleId).
		Where("reply_comment_id IS NOT NULL").Find(&r).Error
	return
}

// FindCommentByArticleAndElected 查询某文章的精选一级评论
func (t *WxCommentDao) FindCommentByArticleAndElected(m *model.WxComment) (r *[]model.WxCommentExt, err error) {
	subQuery := t.db.Model(&model.WxPassport{})
	err = t.db.Model(&model.WxComment{}).
		Select("shippo_wx_comment.*", "temp.nickname AS nickname", "temp.avatar_url AS avatar_url").
		Joins("Left JOIN (?) temp ON temp.id = wx_passport_id", subQuery).
		Order("is_top DESC").Order("like_num DESC").
		Where("article_id", m.ArticleId).
		Where("reply_comment_id IS NULL").
		Where("is_elected", 1).Find(&r).Error
	return
}

// FindReplyByArticleAndElected 查询某文章的精选二级评论
func (t *WxCommentDao) FindReplyByArticleAndElected(m *model.WxComment) (r *[]model.WxCommentExt, err error) {
	subQuery := t.db.Model(&model.WxPassport{})
	err = t.db.Model(&model.WxComment{}).
		Select("shippo_wx_comment.*", "temp.nickname AS nickname", "temp.avatar_url AS avatar_url").
		Joins("Left JOIN (?) temp ON temp.id = wx_passport_id", subQuery).
		Order("created_at ASC").
		Where("article_id", m.ArticleId).
		Where("reply_comment_id IS NOT NULL").
		Where("is_elected", 1).Find(&r).Error
	return
}

// FindCommentByWxPassportAndArticle 查询某用户在某文章的全部一级评论
func (t *WxCommentDao) FindCommentByWxPassportAndArticle(m *model.WxComment) (r *[]model.WxCommentExt, err error) {
	subQuery := t.db.Model(&model.WxPassport{})
	err = t.db.Model(&model.WxComment{}).
		Select("shippo_wx_comment.*", "temp.nickname AS nickname", "temp.avatar_url AS avatar_url").
		Joins("Left JOIN (?) temp ON temp.id = wx_passport_id", subQuery).
		Where("article_id", m.ArticleId).
		Where("wx_passport_id", m.WxPassportId).
		Where("reply_comment_id IS NULL").Find(&r).Error
	return
}

// FindReplyByCommentAndArticle 查询某用户在某文章的全部一级评论 的全部二级评论
func (t *WxCommentDao) FindReplyByCommentAndArticle(m *model.WxComment) (r *[]model.WxCommentExt, err error) {
	subQuery := t.db.Model(&model.WxComment{}).Select("id").Where("article_id", m.ArticleId).
		Where("wx_passport_id", m.WxPassportId).Where("reply_comment_id IS NULL")
	subQuery2 := t.db.Model(&model.WxPassport{})
	err = t.db.Model(&model.WxComment{}).
		Select("shippo_wx_comment.*", "temp.nickname AS nickname", "temp.avatar_url AS avatar_url").
		Joins("Left JOIN (?) temp ON temp.id = wx_passport_id", subQuery2).
		Order("created_at ASC").
		Where("article_id", m.ArticleId).
		Where("reply_comment_id IN (?)", subQuery).Find(&r).Error
	return
}

// FindByWxPassportAndOffiaccount 查询某用户在某公众号的全部一级评论
func (t *WxCommentDao) FindByWxPassportAndOffiaccount(m *model.WxArticle) (r *[]model.WxComment, err error) {
	subQuery := t.db.Model(&model.WxArticle{}).Select("id").Where("offiaccount_id", m.OffiaccountId)
	err = t.db.Order("created_at ASC").Where("article_id IN (?)", subQuery).
		Where("wx_passport_id", m.WxPassportId).Where("reply_comment_id IS NULL").Find(&r).Error
	return
}

// FindByWxPassportAndOffiaccountAndElected 查询某用户在某公众号的精选一级评论
func (t *WxCommentDao) FindByWxPassportAndOffiaccountAndElected(m *model.WxArticle) (r *[]model.WxComment, err error) {
	subQuery := t.db.Model(&model.WxArticle{}).Select("id").Where("offiaccount_id", m.OffiaccountId)
	err = t.db.Order("created_at ASC").Where("article_id IN (?)", subQuery).
		Where("wx_passport_id", m.WxPassportId).Where("reply_comment_id IS NULL").
		Where("is_elected", 1).Find(&r).Error
	return
}

// Find 根据id查询评论
func (t *WxCommentDao) Find(id uint) (r *model.WxComment, err error) {
	err = t.db.First(r, id).Error
	return
}

// UpdateElected 更新评论 精选状态
func (t *WxCommentDao) UpdateElected(m *model.WxComment) (err error) {
	err = t.db.Select("is_elected", "updated_at").Updates(m).Error
	return
}

// UpdateTop 更新评论 置顶状态
func (t *WxCommentDao) UpdateTop(m *model.WxComment) (err error) {
	err = t.db.Select("is_top", "updated_at").Updates(m).Error
	return
}

// UpdateLikeNum 更新评论 点赞数量
func (t *WxCommentDao) UpdateLikeNum(m *model.WxComment) (err error) {
	subQuery := t.db.Model(&model.WxCommentLike{}).Select("COUNT(*)").Where("comment_id", m.ID)
	err = t.db.Model(m).Update("like_num", subQuery).Error
	return
}

// Delete 删除某评论
func (t *WxCommentDao) Delete(id uint) (err error) {
	err = t.db.Delete(&model.WxComment{}, id).Error
	return err
}

// Create 评论
func (t *WxCommentDao) Create(m *model.WxComment) (r *model.WxComment, err error) {
	r = &model.WxComment{
		Content:        m.Content,
		ArticleId:      m.ArticleId,
		WxPassportId:   m.WxPassportId,
		LikeNum:        0,
		IsElected:      0,
		IsTop:          0,
		ReplyCommentId: m.ReplyCommentId,
	}
	if m.ReplyCommentId == 0 {
		err = t.db.Omit("reply_comment_id").Create(r).Error
	} else {
		err = t.db.Create(r).Error
	}
	return
}
