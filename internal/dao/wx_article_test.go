package dao

import (
	"fmt"
	"shippo-server/internal/model"
	"testing"
)

func TestWxArticle_Create(t *testing.T) {
	d := newTest()

	d.Group.WxArticle.Create(&model.WxArticle{
		Title:         "Title",
		Url:           "",
		Image1:        "Image1",
		Image2:        "Image2",
		OffiaccountId: 1,
		WxPassportId:  1,
	})
}

func TestWxArticle_Update(t *testing.T) {
	d := newTest()

	d.Group.WxArticle.Update(&model.WxArticle{
		Model:         model.Model{ID: 1},
		Title:         "Title",
		Url:           "",
		Image1:        "Image1",
		Image2:        "Image2",
		OffiaccountId: 2,
	})
}

func TestWxArticle_UpdateCommentSwitch(t *testing.T) {
	d := newTest()

	d.Group.WxArticle.UpdateCommentSwitch(&model.WxArticle{
		Model:         model.Model{ID: 1},
		CommentSwitch: 0,
	})
}

func TestWxArticle_Find(t *testing.T) {
	d := newTest()

	r, _ := d.Group.WxArticle.Find(1)
	fmt.Printf("TestWxArticle_Find:%+v\n", r)
}
