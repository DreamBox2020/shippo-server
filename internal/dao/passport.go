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

// PassportGet 根据token获取通行证信息
func (t *PassportDao) PassportGet(token string) (p model.Passport, err error) {
	err = t.db.Where("token", token).Limit(1).Find(&p).Error
	return
}

// PassportCreate 创建一个通行证
func (t *PassportDao) PassportCreate(m model.Passport) (r model.Passport, err error) {
	r = model.Passport{
		Token:        utils.GenerateToken(),
		UserId:       m.UserId,
		Ip:           m.Ip,
		Ua:           m.Ua,
		Client:       m.Client,
		WxPassportId: m.WxPassportId,
	}

	tx := t.db

	if m.UserId == 0 {
		tx = tx.Omit("user_id")
	}

	if m.WxPassportId == 0 {
		tx = tx.Omit("wx_passport_id")
	}

	err = tx.Create(&r).Error

	return

}

// PassportDelete 根据uid删除该用户全部客户端的通行证
func (t *PassportDao) PassportDelete(userId uint, client uint) error {
	return t.db.Where("user_id", userId).Where("client", client).Delete(&model.Passport{}).Error
}

// PassportUpdate 根据token更新通行证信息
func (t *PassportDao) PassportUpdate(token string, m model.Passport) (err error) {

	var selects = []string{
		"updated_at",
	}

	if m.UserId != 0 {
		selects = append(selects, "user_id")
	}

	if m.Ip != "" {
		selects = append(selects, "ip")
	}

	if m.Ua != "" {
		selects = append(selects, "ua")
	}

	err = t.db.Select(selects).Where("token", token).Updates(&m).Error

	return
}
