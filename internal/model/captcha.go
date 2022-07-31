package model

import "time"

const (
	CaptchaTypePhone = 0
	CaptchaTypeEmail = 1
)

type Captcha struct {
	Model
	Target string
	Code   string
	Token  string
	Type   int
}

// IsExpire 是否失效 有效期15分钟
func (t *Captcha) IsExpire() bool {
	return time.Since(t.UpdatedAt) > time.Minute*15
}
