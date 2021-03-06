package dao

import "shippo-server/internal/model"

func (d *Dao) UserFindByPhone(phone string) (u model.User, err error) {
	err = d.db.Where("phone", phone).Limit(1).Find(&u).Error
	return
}

func (d *Dao) UserCreate(phone string) (u model.User, err error) {
	u.Phone = phone
	err = d.db.Omit("nickname", "email", "avatar").Create(&u).Error
	return
}
