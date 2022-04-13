package model

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
