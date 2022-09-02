package dao

import (
	"shippo-server/internal/model"
	"shippo-server/utils"
)

type CaptchaDao struct {
	*Dao
}

func NewCaptchaDao(s *Dao) *CaptchaDao {
	return &CaptchaDao{s}
}

func (d *CaptchaDao) CaptchaSmsInsert(target string, token string, channel string) (s model.Captcha, err error) {
	s.Target = target
	s.Code = utils.GenerateCaptcha()
	s.Token = token
	s.Type = model.CaptchaTypePhone
	s.Channel = channel
	err = d.db.Create(&s).Error
	return
}

func (d *CaptchaDao) CaptchaEmailInsert(target string, token string, channel string) (s model.Captcha, err error) {
	s.Target = target
	s.Code = utils.GenerateCaptcha()
	s.Token = token
	s.Type = model.CaptchaTypeEmail
	s.Channel = channel
	err = d.db.Create(&s).Error
	return
}

func (d *CaptchaDao) CaptchaDel(target string) error {
	return d.db.Where("target = ?", target).Delete(&model.Captcha{}).Error
}

// FindByTargetAndCode 查询验证码
func (d *CaptchaDao) FindByTargetAndCode(m *model.Captcha) (s model.Captcha, err error) {
	err = d.db.Where("target", m.Target).
		Where("code", m.Code).
		Where("token", m.Token).
		Where("channel", m.Channel).
		First(&s).Error
	return
}

// FindShadowByTargetAndCode 查询已经被删除的验证码
func (d *CaptchaDao) FindShadowByTargetAndCode(m *model.Captcha) (r model.Captcha, err error) {
	err = d.db.Unscoped().
		Where("target", m.Target).
		Where("code", m.Code).
		Where("token", m.Token).
		Where("channel", m.Channel).
		Where("deleted_at IS NOT NULL").
		First(&r).Error
	return
}
