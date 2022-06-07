package dao

import (
	"shippo-server/internal/model"
)

type WxCommentLikeDao struct {
	*Dao
}

func NewWxCommentLikeDao(s *Dao) *WxCommentLikeDao {
	return &WxCommentLikeDao{s}
}

// Create 新增点赞
func (t *WxCommentLikeDao) Create(m *model.WxCommentLike) (r *model.WxCommentLike, err error) {
	r = &model.WxCommentLike{
		CommentId:    m.CommentId,
		WxPassportId: m.WxPassportId,
	}
	err = t.db.Create(r).Error
	return
}

// Delete 取消点赞
func (t *WxCommentLikeDao) Delete(m *model.WxCommentLike) (err error) {
	err = t.db.Where("comment_id", m.CommentId).Where("wx_passport_id", m.WxPassportId).Delete(&model.WxCommentLike{}).Error
	return
}

// Find 查询点赞
func (t *WxCommentLikeDao) Find(commentId uint, wxPassportId uint) (r *model.WxCommentLike, err error) {
	err = t.db.Where("comment_id", commentId).Where("wx_passport_id", wxPassportId).First(r).Error
	return
}
