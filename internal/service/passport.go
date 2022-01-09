package service

import (
	"fmt"
	"shippo-server/internal/model"
	"shippo-server/utils/box"
	"shippo-server/utils/check"
	"shippo-server/utils/ecode"
	"time"
)

func (s *Service) PassportCreate(c *box.Context, passport string, ip string) (data map[string]interface{}, err error) {
	fmt.Printf("service->PassportCreate->args->passport:%+v\n", passport)
	fmt.Printf("service->PassportCreate->args->ip:%+v\n", ip)

	p, err := s.PassportGet(c, passport, ip)

	if passport != "" && err != nil {
		return
	}

	// 如果不存在或者到期(30天)，就创建一个新的通行证，否则，就续期旧的。
	if p.Token == "" || time.Since(p.UpdatedAt) > time.Hour*24*30 {
		p = model.Passport{
			Token:  "",
			UserId: 0,
			Ip:     ip,
			Ua:     "",
			Client: 0,
		}

		p, err = s.dao.PassportCreate(p)
		if err != nil {
			return
		}

	} else {

		p, err = s.dao.PassportUpdate(p.Token, model.Passport{Ip: ip})
		if err != nil {
			return
		}
	}

	data = make(map[string]interface{})
	data["passport"] = p.Token
	data["uid"] = p.UserId

	return
}

func (s *Service) PassportGet(c *box.Context, passport string, ip string) (p model.Passport, err error) {
	if !check.CheckPassport(passport) {
		err = ecode.ServerErr
		return
	}
	return s.dao.PassportGet(passport)
}
