package model

import "gorm.io/gorm"

type Sms struct {
	gorm.Model
	Phone string
	Code  string
	Token string
}
