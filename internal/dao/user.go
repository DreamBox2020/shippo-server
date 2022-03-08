package dao

import "shippo-server/internal/model"

type UserDao struct {
	*Dao
}

func NewUserDao(s *Dao) *UserDao {
	return &UserDao{s}
}

func (d *UserDao) UserFindByUID(uid uint) (u model.User, err error) {
	err = d.db.Where("id", uid).Limit(1).Find(&u).Error
	return
}

func (d *UserDao) UserFindByPhone(phone string) (u model.User, err error) {
	err = d.db.Where("phone", phone).Limit(1).Find(&u).Error
	return
}

func (d *UserDao) UserFindByEmail(email string) (u model.User, err error) {
	err = d.db.Where("email", email).Limit(1).Find(&u).Error
	return
}

func (d *UserDao) UserCreate(phone string) (u model.User, err error) {
	u.Phone = phone
	err = d.db.Omit("nickname", "email", "avatar").Create(&u).Error
	return
}

func (d *UserDao) UserCreateEmail(email string) (u model.User, err error) {
	u.Email = email
	err = d.db.Omit("nickname", "phone", "avatar").Create(&u).Error
	return
}
