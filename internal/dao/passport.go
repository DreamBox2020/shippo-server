package dao

import (
	"github.com/satori/go.uuid"
	"shippo-server/internal/model"
	"strings"
)

// 根据token获取通行证信息
func (d *Dao) PassportGet(token string) (p model.Passport, err error) {
	err = d.db.Where("token", token).Limit(1).Find(&p).Error
	return
}

// 创建一个通行证
func (d *Dao) PassportCreate(p model.Passport) (model.Passport, error) {
	// 生成token
	p.Token = strings.Replace(uuid.NewV4().String(), "-", "", -1)
	return p, d.db.Create(&p).Error
}

// 根据uid删除该用户全部客户端的通行证
func (d *Dao) PassportDelete(userId int, client int) error {
	return d.db.Where("user_id", userId).Where("client", client).Delete(model.Passport{}).Error
}

// 根据token更新通行证信息
func (d *Dao) PassportUpdate(token string, p model.Passport) (res model.Passport, err error) {
	err = d.db.Model(&model.Passport{}).Where("token", token).Updates(&p).Error

	if err != nil {
		return
	}

	res, err = d.PassportGet(token)
	return
}
