package model

import "gorm.io/gorm"

const (
	CaptchaTypePhone = 0
	CaptchaTypeEmail = 1
)

type Captcha struct {
	gorm.Model
	Target string
	Code   string
	Token  string
	Type   int
}
