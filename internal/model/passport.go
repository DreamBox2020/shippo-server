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
