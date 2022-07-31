package model

import "time"

const (
	CaptchaTypePhone = 0
	CaptchaTypeEmail = 1
)

type Captcha struct {
	Model
	Target  string
	Code    string
	Token   string
	Type    int // 0: phone, 1: email
	Channel string
}

// IsExpire 是否失效 有效期15分钟 true为失效
func (t *Captcha) IsExpire() bool {
	return time.Since(t.UpdatedAt) > time.Minute*15
}

// ShadowIsExpire 影子是否失效 有效期30分钟 true为失效
func (t *Captcha) ShadowIsExpire(c *Captcha) bool {
	// 影子必须已经被删除
	if t.DeletedAt.Valid {
		return time.Since(c.DeletedAt.Time) > time.Minute*30
	}

	return false
}
