package dao

import "shippo-server/internal/model"

type UserDao struct {
	*Dao
}

func NewUserDao(s *Dao) *UserDao {
	return &UserDao{s}
}

func (t *UserDao) UserFindByUID(uid uint) (u model.User, err error) {
	err = t.db.Where("id", uid).Limit(1).Find(&u).Error
	return
}

func (t *UserDao) UserFindByPhone(phone string) (u model.User, err error) {
	err = t.db.Where("phone", phone).Limit(1).Find(&u).Error
	return
}

func (t *UserDao) UserFindByEmail(email string) (u model.User, err error) {
	err = t.db.Where("email", email).Limit(1).Find(&u).Error
	return
}

func (t *UserDao) UserCreate(phone string) (u model.User, err error) {
	u.Phone = phone
	u.Role = 2
	err = t.db.Omit("email", "nickname", "avatar", "wx_passport_id").Create(&u).Error
	return
}

func (t *UserDao) UserCreateEmail(email string) (u model.User, err error) {
	u.Email = email
	u.Role = 2
	err = t.db.Omit("phone", "nickname", "avatar", "wx_passport_id").Create(&u).Error
	return
}

func (t *UserDao) FindAll(u model.UserFindAllReq) (m model.UserFindAllResp, err error) {

	m.Pagination.Copy(u.Pagination)

	db := t.db.Model(&model.User{})

	if u.ID != 0 {
		db = db.Where("id", u.ID)
	}

	if u.Phone != "" {
		db = db.Where("phone", u.Phone)
	}

	if u.Email != "" {
		db = db.Where("email", u.ID)
	}

	if u.Nickname != "" {
		db = db.Where("nickname", u.Nickname)
	}

	err = db.Count(&m.Total).Error
	if err != nil {
		return
	}

	subQuery := t.db.Model(&model.Role{}).Select("id", "role_name")
	db = db.Select("shippo_user.*", "temp.role_name").
		Joins("Left JOIN (?) temp ON temp.id = role", subQuery)

	err = db.Scopes(u.Pagination.Sql()).Find(&m.Items).Error

	return
}

func (t *UserDao) UpdateUserRole(u model.User) error {
	return t.db.Model(&model.User{}).Where("id", u.ID).
		Update("role", u.Role).Error
}
