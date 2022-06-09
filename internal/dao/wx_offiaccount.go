package dao

import (
	"shippo-server/internal/model"
)

type WxOffiaccountDao struct {
	*Dao
}

func NewWxOffiaccountDao(s *Dao) *WxOffiaccountDao {
	return &WxOffiaccountDao{s}
}

// FindAll 查询所有公众号
func (t *WxOffiaccountDao) FindAll() (r *[]model.WxOffiaccount, err error) {
	err = t.db.Find(&r).Error
	return
}

// FindByUsername 根据username查询公众号
func (t *WxOffiaccountDao) FindByUsername(username string) (u *model.WxOffiaccount, err error) {
	err = t.db.Where("username", username).First(&u).Error
	return
}

// Find 根据id查询公众号
func (t *WxOffiaccountDao) Find(id uint) (u *model.WxOffiaccount, err error) {
	err = t.db.First(&u, id).Error
	return
}
