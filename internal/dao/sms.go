package dao

import (
	"math/rand"
	"shippo-server/internal/model"
	"strconv"
	"time"
)

// 生成一个记录，验证码随机生成
func (d *Dao) SmsInsert(phone string, token string) (s model.Sms, err error) {
	rand.Seed(time.Now().UnixNano())
	code := strconv.Itoa(rand.Intn(899999) + 100000)

	s.Phone = phone
	s.Code = code
	s.Token = token
	err = d.db.Create(&s).Error
	return
}

// 根据手机号删除一个记录
func (d *Dao) SmsDel(phone string) error {
	return d.db.Where("phone = ?", phone).Delete(&model.Sms{}).Error
}

// 根据手机号和验证码查询一个记录
func (d *Dao) SmsByPhoneAndCode(phone string, code string, token string) (s model.Sms, err error) {
	err = d.db.Where("phone", phone).Where("code", code).Where("token", token).Limit(1).Find(&s).Error
	return
}
