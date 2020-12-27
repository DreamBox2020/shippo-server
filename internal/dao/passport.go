package dao

import (
	"shippo-server/internal/model"
)

func (d *Dao) GetPassport(passport string) (p model.Passport, err error) {
	p = model.Passport{
		Token: passport,
	}
	err = d.db.First(&p).Error
	return
}
