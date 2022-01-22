package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Phone    string
	Email    string
	Nickname string
	Avatar   string
	Exp      uint
	Coin     uint
	Role     uint
}

type UserLoginParam struct {
	Phone string `json:"phone"`
	Email string `json:"email"`
	Code  string `json:"code"`
}
