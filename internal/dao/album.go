package dao

import (
	"shippo-server/internal/model"
	"shippo-server/utils/ecode"
)

// 查询一个用户所拥有的全部相簿
func (d *Dao) AlbumFindAll(uid uint) (p []model.Album, err error) {
	err = d.db.Where("user_id", uid).Find(&p).Error
	return
}

// 查询一个用户所拥有的名叫？的相簿
func (d *Dao) AlbumFind(uid uint, name string) (p model.Album, err error) {
	err = d.db.Where("user_id", uid).Where("name", name).Limit(1).Find(&p).Error
	return
}

// 查询一个用户是否拥有的名叫？的相簿
func (d *Dao) AlbumHas(uid uint, name string) bool {
	var count int64
	if err := d.db.Model(&model.Album{}).Where("user_id", uid).Where("name", name).Limit(1).Count(&count).Error; err != nil {
		// 如果出错，就判断为 存在
		return true
	}
	return count > 0
}

// 创建一个相簿
func (d *Dao) AlbumCreate(p model.Album) (err error) {
	if d.AlbumHas(p.UserId, p.Name) {
		err = ecode.ServerErr
		return
	}

	err = d.db.Create(p).Error
	return
}

// 删除一个相簿
func (d *Dao) AlbumDelete(id uint) (err error) {
	err = d.db.Where("id", id).Delete(&model.Album{}).Error
	return
}

// 更新一个相簿的数据
func (d *Dao) AlbumUpdate(p model.Album) (err error) {
	err = d.db.Model(&model.Album{}).Select("name", "intro").Updates(&p).Error
	return
}
