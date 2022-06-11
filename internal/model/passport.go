package model

import "time"

type Passport struct {
	Model
	Token        string
	UserId       uint
	Ip           string
	Ua           string
	Client       uint
	WxPassportId uint // 如果是从微信渠道进入的，则该字段有值。
}

// IsLogin 是否登录
func (t *Passport) IsLogin() bool {
	return t.UserId != 0
}

// IsExpire 是否失效 有效期30天
func (t *Passport) IsExpire() bool {
	return time.Since(t.UpdatedAt) > time.Hour*24*30
}

type PassportCreateResult struct {
	Passport string             `json:"passport"`
	Uid      uint               `json:"uid"`
	Access   []PermissionAccess `json:"access"`
}
