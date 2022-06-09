package service

import (
	"shippo-server/internal/model"
)

type WxCommentService struct {
	*Service
}

func NewWxCommentService(s *Service) *WxCommentService {
	return &WxCommentService{s}
}

// FindByArticle 查询某文章精选评论
func (t *WxCommentService) FindByArticle(m *model.WxComment) (r *[]model.WxCommentExtReplyList, err error) {
	// 查精选一级评论
	comments, err := t.dao.WxComment.FindCommentByArticleAndElected(m)
	if err != nil {
		return
	}

	// 查精选二级评论
	replies, err := t.dao.WxComment.FindReplyByArticleAndElected(m)
	if err != nil {
		return
	}

	r = t.mergeComment(comments, replies)

	return
}

// AdminFindByArticle 查询某文章全部评论
func (t *WxCommentService) AdminFindByArticle(m *model.WxComment) (r *[]model.WxCommentExtReplyList, err error) {
	// 查全部一级评论
	comments, err := t.dao.WxComment.FindCommentByArticle(m)
	if err != nil {
		return
	}

	// 查全部二级评论
	replies, err := t.dao.WxComment.FindReplyByArticle(m)
	if err != nil {
		return
	}

	r = t.mergeComment(comments, replies)

	return
}

// mergeComment 合并评论
func (t *WxCommentService) mergeComment(comments, replies *[]model.WxComment) *[]model.WxCommentExtReplyList {
	// 转为map结构
	var commentMap = make(map[uint]model.WxCommentExtReplyList, len(*comments))
	for _, comment := range *comments {
		commentMap[comment.ID] = model.WxCommentExtReplyList{
			WxComment: comment,
		}
	}

	// 将二级评论添加到一级评论的 ReplyList 属性中
	for _, reply := range *replies {
		comment, ok := commentMap[reply.ReplyCommentId]
		if ok {
			comment.ReplyList = append(comment.ReplyList, reply)
			commentMap[reply.ReplyCommentId] = comment
		}
	}

	// 转为slice结构
	var result = make([]model.WxCommentExtReplyList, len(commentMap))
	for _, comment := range commentMap {
		result = append(result, comment)
	}

	return &result
}

// FindByWxPassportAndArticle 查询当前⽤户在某⽂章的全部评论
func (t *WxCommentService) FindByWxPassportAndArticle(m *model.WxComment) (r *[]model.WxCommentExtReplyList, err error) {
	// 查全部一级评论
	comments, err := t.dao.WxComment.FindCommentByWxPassportAndArticle(m)
	if err != nil {
		return
	}

	// 查全部二级评论
	replies, err := t.dao.WxComment.FindReplyByCommentAndArticle(m)
	if err != nil {
		return
	}

	r = t.mergeComment(comments, replies)

	return
}

// FindByWxPassportAndOffiaccount 查询某用户在某公众号的全部一级评论
func (t *WxCommentService) FindByWxPassportAndOffiaccount(m *model.WxArticle) (r *[]model.WxComment, err error) {
	r, err = t.dao.WxComment.FindByWxPassportAndOffiaccount(m)
	return
}

// FindByWxPassportAndOffiaccountAndElected 查询某用户在某公众号的精选一级评论
func (t *WxCommentService) FindByWxPassportAndOffiaccountAndElected(m *model.WxArticle) (r *[]model.WxComment, err error) {
	r, err = t.dao.WxComment.FindByWxPassportAndOffiaccountAndElected(m)
	return
}

// UpdateElected 更新评论 精选状态
func (t *WxCommentService) UpdateElected(m *model.WxComment) (err error) {
	err = t.dao.WxComment.UpdateElected(m)
	return
}

// UpdateTop 更新评论 置顶状态
func (t *WxCommentService) UpdateTop(m *model.WxComment) (err error) {
	err = t.dao.WxComment.UpdateTop(m)
	return
}

// UpdateLikeNum 更新评论 点赞数量
func (t *WxCommentService) UpdateLikeNum(m *model.WxComment) (err error) {
	err = t.dao.WxComment.UpdateLikeNum(m)
	return
}

// Delete 删除某评论
func (t *WxCommentService) Delete(m *model.WxComment) (err error) {
	err = t.dao.WxComment.Delete(m.ID)
	return
}

// Create 评论
func (t *WxCommentService) Create(m *model.WxComment) (r *model.WxComment, err error) {
	m.ReplyCommentId = 0
	r, err = t.dao.WxComment.Create(m)
	return
}

// Reply 回复评论
func (t *WxCommentService) Reply(m *model.WxComment) (r *model.WxComment, err error) {
	r, err = t.dao.WxComment.Create(m)
	return
}

// AdminReply 管理员回复评论
func (t *WxCommentService) AdminReply(m *model.WxComment) (r *model.WxComment, err error) {
	r, err = t.dao.WxComment.Create(m)
	return
}
