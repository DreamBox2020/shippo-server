package dao

import "shippo-server/internal/model"

func (d *Dao) UserFindByPhone(phone string) (u model.User, err error) {
	err = d.db.Where("phone", phone).Limit(1).Find(&u).Error
	return
}

func (d *Dao) UserFindByEmail(email string) (u model.User, err error) {
	err = d.db.Where("email", email).Limit(1).Find(&u).Error
	return
}

func (d *Dao) UserCreate(phone string) (u model.User, err error) {
	u.Phone = phone
	err = d.db.Omit("nickname", "email", "avatar").Create(&u).Error
	return
}

func (d *Dao) UserCreateEmail(email string) (u model.User, err error) {
	u.Email = email
	err = d.db.Omit("nickname", "phone", "avatar").Create(&u).Error
	return
}
