package dao

import (
	"shippo-server/internal/model"
	"shippo-server/utils"
)

type PassportDao struct {
	*Dao
}

func NewPassportDao(s *Dao) *PassportDao {
	return &PassportDao{s}
}

// 根据token获取通行证信息
func (d *PassportDao) PassportGet(token string) (p model.Passport, err error) {
	err = d.db.Where("token", token).Limit(1).Find(&p).Error
	return
}

// 创建一个通行证
func (d *PassportDao) PassportCreate(p model.Passport) (model.Passport, error) {
	// 生成token
	p.Token = utils.GenerateToken()
	return p, d.db.Create(&p).Error
}

// 根据uid删除该用户全部客户端的通行证
func (d *PassportDao) PassportDelete(userId uint, client uint) error {
	return d.db.Where("user_id", userId).Where("client", client).Delete(model.Passport{}).Error
}

// 根据token更新通行证信息
func (d *PassportDao) PassportUpdate(token string, p model.Passport) (res model.Passport, err error) {
	err = d.db.Model(&model.Passport{}).Where("token", token).Updates(&p).Error

	if err != nil {
		return
	}

	res, err = d.PassportGet(token)
	return
}
