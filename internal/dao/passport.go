package dao

import (
	"github.com/satori/go.uuid"
	"shippo-server/internal/model"
	"strings"
)

func (d *Dao) GetPassport(passport string) (p model.Passport, err error) {
	p.Token = passport
	err = d.db.First(&p).Error
	return
}

func (d *Dao) CreatePassport() (p model.Passport, err error) {
	p.Token = strings.Replace(uuid.NewV4().String(), "-", "", -1)
	err = d.db.Create(&p).Error
	return
}
