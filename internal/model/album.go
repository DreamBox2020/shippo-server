package model

import "gorm.io/gorm"

type Album struct {
	gorm.Model
	Name   string
	Intro  string
	UserId uint
}
