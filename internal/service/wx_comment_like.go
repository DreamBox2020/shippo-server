package service

import (
	"errors"
	"gorm.io/gorm"
	"shippo-server/internal/model"
)

type WxCommentLikeService struct {
	*Service
}

func NewWxCommentLikeService(s *Service) *WxCommentLikeService {
	return &WxCommentLikeService{s}
}

// Create 新增点赞
func (t *WxCommentLikeService) Create(m *model.WxCommentLike) (r *model.WxCommentLike, err error) {

	//⾸先查询⼀遍，有没有点过赞，点过，就不继续点赞。
	c, err := t.dao.WxCommentLike.Find(m.CommentId, m.WxPassportId)
	if err != nil {
		// 如果没有查到
		if errors.Is(err, gorm.ErrRecordNotFound) {
			r, err = t.dao.WxCommentLike.Create(m)
		}
		return
	}
	return c, err
}

// Delete 取消点赞
func (t *WxCommentLikeService) Delete(m *model.WxCommentLike) (err error) {
	err = t.dao.WxCommentLike.Delete(m)
	return
}

// Find 查询点赞
func (t *WxCommentLikeService) Find(commentId uint, wxPassportId uint) (r *model.WxCommentLike, err error) {
	r, err = t.dao.WxCommentLike.Find(commentId, wxPassportId)
	return
}
