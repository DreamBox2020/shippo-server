package dao

import (
	"shippo-server/internal/model"
	"shippo-server/utils"
)

func (d *Dao) CaptchaSmsInsert(target string, token string) (s model.Captcha, err error) {
	s.Target = target
	s.Code = utils.GenerateCaptcha()
	s.Token = token
	s.Type = model.CaptchaTypePhone
	err = d.db.Create(&s).Error
	return
}

func (d *Dao) CaptchaEmailInsert(target string, token string) (s model.Captcha, err error) {
	s.Target = target
	s.Code = utils.GenerateCaptcha()
	s.Token = token
	s.Type = model.CaptchaTypeEmail
	err = d.db.Create(&s).Error
	return
}

func (d *Dao) CaptchaDel(target string) error {
	return d.db.Where("target = ?", target).Delete(&model.Captcha{}).Error
}

func (d *Dao) CaptchaByTargetAndCode(target string, code string, token string) (s model.Captcha, err error) {
	err = d.db.Where("target", target).Where("code", code).Where("token", token).Limit(1).Find(&s).Error
	return
}
