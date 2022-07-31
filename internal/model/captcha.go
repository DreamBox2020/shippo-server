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

// ShadowIsExpire 影子是否失效 有效期30分钟
func (t *Captcha) ShadowIsExpire(c *Captcha) bool {
	// 如果影子被删除的时间，大于30分钟，则无效
	if time.Since(c.DeletedAt.Time) > time.Minute*30 {
		return true
	}

	return t.IsExpire()
}
