package service

import (
	"errors"
	"gorm.io/gorm"
	"shippo-server/internal/model"
	"shippo-server/utils/ecode"
	"shippo-server/utils/wx"
)

type WxArticleService struct {
	*Service
}

func NewWxArticleService(s *Service) *WxArticleService {
	return &WxArticleService{s}
}

// Create 新增文章
func (t *WxArticleService) Create(m *model.WxArticle) (r *model.WxArticle, err error) {
	article, err := t.createWxArticle(m)

	r, err = t.dao.WxArticle.Create(article)
	return
}

// Update 修改文章
func (t *WxArticleService) Update(m *model.WxArticle) (err error) {

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

	article, err := t.createWxArticle(m)

	err = t.dao.WxArticle.Update(article)
	return
}

// createWxArticle 创建一个文章模型
func (t *WxArticleService) createWxArticle(m *model.WxArticle) (r *model.WxArticle, err error) {

	// 获取文章模型
	article, err := wx.NewArticle(m.Url)
	if err != nil {
		return
	}

	if !article.IsWX() {
		err = ecode.WxArticleURLError
		return
	}

	// 获取公众号数据
	offiaccount, err := t.dao.WxOffiaccount.FindByUsername(article.Username())
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = ecode.WxOffiaccountIsNotLinked
		}
		return
	}

	image1, err := t.Group.File.ToLocalUrl(article.Image1(), "wx")
	if err != nil {
		return
	}

	image2, err := t.Group.File.ToLocalUrl(article.Image2(), "wx")
	if err != nil {
		return
	}

	r = &model.WxArticle{
		Title:         article.Title(),
		Image1:        image1.Uri,
		Image2:        image2.Uri,
		OffiaccountId: offiaccount.ID,
		WxPassportId:  m.WxPassportId,
	}

	if !article.IsTempURL() {
		r.Url = article.URL()
	}

	if m.ID != 0 {
		r.ID = m.ID
	}

	return
}

// UpdateCommentSwitch 修改文章评论开关
func (t *WxArticleService) UpdateCommentSwitch(m *model.WxArticle) (err error) {
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
func (t *WxArticleService) FindByOffiaccount(m *model.WxArticle) (r *[]model.WxArticle, err error) {
	r, err = t.dao.WxArticle.FindByOffiaccount(m)
	return
}

// Find 查询文章根据id
func (t *WxArticleService) Find(id uint) (r *model.WxArticleExtOffiaccountNickname, err error) {
	r, err = t.dao.WxArticle.Find(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = ecode.ErrRecordNotFound
	}
	return
}

// FindAllByWxPassport 查询某人的全部文章
func (t *WxArticleService) FindAllByWxPassport(m *model.WxArticle) (r *[]model.WxArticleExtOffiaccountNickname, err error) {
	r, err = t.dao.WxArticle.FindAllByWxPassport(m)
	return
}
