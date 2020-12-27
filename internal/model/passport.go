package model

import "gorm.io/gorm"

type Passport struct {
	gorm.Model
	Token  string
	UserId int
	Ip     string
	Ua     string
	Client int
}
