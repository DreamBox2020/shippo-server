package dao

import (
	"math/rand"
	"shippo-server/internal/model"
	"strconv"
	"time"
)

// 生成一个记录，验证码随机生成
func (d *Dao) SmsInsert(phone string) (s model.Sms, err error) {
	rand.Seed(time.Now().UnixNano())
	code := strconv.Itoa(rand.Intn(999999))

	s.Phone = phone
	s.Code = code
	err = d.db.Create(&s).Error
	return
}

// 根据手机号删除一个记录
func (d *Dao) SmsDel(phone string) error {
	return d.db.Where("phone = ?", phone).Delete(&model.Sms{}).Error
}

// 根据手机号和验证码查询一个记录
func (d *Dao) SmsByPhoneAndCode(phone string, code string) (s model.Sms, err error) {
	s.Phone = phone
	s.Code = code

	err = d.db.First(&s).Error
	return
}
