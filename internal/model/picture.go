package model

import "gorm.io/gorm"

type Picture struct {
	gorm.Model
	Path string
	Uri  string
	Name string
	Mime string
	Type string
}
