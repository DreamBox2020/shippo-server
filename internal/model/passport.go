package model

import "gorm.io/gorm"

type Passport struct {
	gorm.Model
	Token  string
	UserId uint
	Ip     string
	Ua     string
	Client uint
}

func (t *Passport) IsLogin() bool {
	return t.UserId != 0
}

type PassportCreateResult struct {
	Passport string `json:"passport"`
	Uid      uint   `json:"uid"`
}
