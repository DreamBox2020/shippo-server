package service

import (
	"errors"
	"gorm.io/gorm"
	"shippo-server/internal/model"
	"shippo-server/utils/ecode"
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
func (t *WxCommentService) mergeComment(comments, replies *[]model.WxCommentExt) *[]model.WxCommentExtReplyList {
	// 转为map结构
	var commentMap = make(map[uint]model.WxCommentExtReplyList, len(*comments))
	for _, comment := range *comments {
		commentMap[comment.ID] = model.WxCommentExtReplyList{
			WxCommentExt: comment,
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
	var result []model.WxCommentExtReplyList
	for _, comment := range *comments {
		c := commentMap[comment.ID]
		if len(c.ReplyList) == 0 {
			c.ReplyList = make([]model.WxCommentExt, 0)
		}
		result = append(result, c)
	}

	if len(result) == 0 {
		result = make([]model.WxCommentExtReplyList, 0)
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

	isAdmin, err := t.IsAdmin(m.ID, m.WxPassportId)
	if err != nil {
		return
	}

	if !isAdmin {
		err = ecode.ServerErr
		return
	}

	err = t.dao.WxComment.UpdateElected(m)
	return
}

// UpdateTop 更新评论 置顶状态
func (t *WxCommentService) UpdateTop(m *model.WxComment) (err error) {

	isAdmin, err := t.IsAdmin(m.ID, m.WxPassportId)
	if err != nil {
		return
	}

	if !isAdmin {
		err = ecode.ServerErr
		return
	}

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
	// 查询要删除的评论，是否存在
	comment, err := t.dao.WxComment.Find(m.ID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = ecode.ServerErr
		}
		return
	}

	// 查询评论的文章，是否存在
	article, err := t.dao.WxArticle.Find(comment.ArticleId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = ecode.ServerErr
		}
		return
	}

	// 只能删除自己的评论
	if comment.WxPassportId == 0 {
		// 如果不是管理员
		if article.WxPassportId != m.WxPassportId {
			err = ecode.ServerErr
			return
		}
	} else {
		if comment.WxPassportId != m.WxPassportId {
			err = ecode.ServerErr
			return
		}
	}

	err = t.dao.WxComment.Delete(m.ID)
	return
}

// Create 评论
func (t *WxCommentService) Create(m *model.WxComment) (r *model.WxComment, err error) {

	// 评论内容不能为空
	if m.Content == "" {
		err = ecode.WxCommentIsEmptyErr
		return
	}

	// 查询要评论的文章，是否存在
	_, err = t.dao.WxArticle.Find(m.ArticleId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = ecode.ServerErr
		}
		return
	}

	r, err = t.dao.WxComment.Create(&model.WxComment{
		Content:        m.Content,
		ArticleId:      m.ArticleId,
		WxPassportId:   m.WxPassportId,
		LikeNum:        0,
		IsElected:      0,
		IsTop:          0,
		ReplyCommentId: 0,
	})
	return
}

// Reply 回复评论
func (t *WxCommentService) Reply(m *model.WxComment) (r *model.WxComment, err error) {

	// 评论内容不能为空
	if m.Content == "" {
		err = ecode.WxCommentIsEmptyErr
		return
	}

	// 查询要回复的评论，是否存在
	comment, err := t.dao.WxComment.Find(m.ReplyCommentId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = ecode.ServerErr
		}
		return
	}

	// 只能回复自己的评论
	if comment.WxPassportId != m.WxPassportId {
		err = ecode.WxCommentErr
		return
	}

	r, err = t.dao.WxComment.Create(&model.WxComment{
		Content:        m.Content,
		ArticleId:      comment.ArticleId,
		WxPassportId:   m.WxPassportId,
		LikeNum:        0,
		IsElected:      0,
		IsTop:          0,
		ReplyCommentId: m.ReplyCommentId,
	})
	return
}

// AdminReply 管理员回复评论
func (t *WxCommentService) AdminReply(m *model.WxComment) (r *model.WxComment, err error) {

	// 评论内容不能为空
	if m.Content == "" {
		err = ecode.WxCommentIsEmptyErr
		return
	}

	// 查询要回复的评论，是否存在
	comment, err := t.dao.WxComment.Find(m.ReplyCommentId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = ecode.ServerErr
		}
		return
	}

	// 查询要评论的文章，是否存在
	article, err := t.dao.WxArticle.Find(comment.ArticleId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = ecode.ServerErr
		}
		return
	}

	// 管理员只能回复自己文章上的评论
	if article.WxPassportId != m.WxPassportId {
		err = ecode.WxCommentErr
		return
	}

	r, err = t.dao.WxComment.Create(&model.WxComment{
		Content:        m.Content,
		ArticleId:      comment.ArticleId,
		WxPassportId:   0,
		LikeNum:        0,
		IsElected:      0,
		IsTop:          0,
		ReplyCommentId: m.ReplyCommentId,
	})

	return
}

// IsAdmin 检查通行证是否是某评论文章的作者
func (t *WxCommentService) IsAdmin(commentId uint, wxPassportId uint) (r bool, err error) {

	// 查询评论，是否存在
	comment, err := t.dao.WxComment.Find(commentId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = ecode.ServerErr
		}
		return
	}

	// 查询评论的文章，是否存在
	article, err := t.dao.WxArticle.Find(comment.ArticleId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = ecode.ServerErr
		}
		return
	}

	// 是否管理员
	r = article.WxPassportId == wxPassportId
	return
}
