package dao

import "shippo-server/internal/model"

type PictureDao struct {
	*Dao
}

func NewPictureDao(s *Dao) *PictureDao {
	return &PictureDao{s}
}

// 创建一条记录
func (t *PictureDao) Create(m model.Picture) (r model.Picture, err error) {
	err = t.db.Create(&m).Error
	return m, err
}

// 删除一条记录
func (t *PictureDao) Delete(id uint) (err error) {
	err = t.db.Where("id", id).Delete(&model.Picture{}).Error
	return
}

// 根据uri获取一条记录
func (t *PictureDao) FindByUri(uri string) (r model.Picture, err error) {
	err = t.db.Where("uri", uri).Limit(1).Find(&r).Error
	return
}
