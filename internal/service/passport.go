package service

import (
	"fmt"
	"shippo-server/internal/model"
	"shippo-server/utils/box"
	"shippo-server/utils/check"
	"shippo-server/utils/ecode"
)

func (s *Service) PassportCreate(c *box.Context, passport string, ip string) (interface{}, error) {
	fmt.Printf("service->PassportCreate->args->passport:%+v\n", passport)
	fmt.Printf("service->PassportCreate->args->ip:%+v\n", ip)

	p, err := s.dao.PassportGet(passport)

	if err != nil {
		return nil, err
	}

	// 如果不存在，就创建一个新的通行证，否则，就续期旧的。
	if p.Token == "" {
		p = model.Passport{
			Token:  "",
			UserId: 0,
			Ip:     ip,
			Ua:     "",
			Client: 0,
		}

		p, err = s.dao.PassportCreate(p)
		if err != nil {
			return nil, err
		}

	} else {

		p, err = s.dao.PassportUpdate(p.Token, model.Passport{Ip: ip})
		if err != nil {
			return nil, err
		}
	}

	data := make(map[string]interface{}, 2)
	data["passport"] = p.Token
	data["uid"] = p.UserId

	return data, err
}

func (s *Service) PassportGet(c *box.Context, passport string, ip string) (p model.Passport, err error) {
	if !check.CheckPassport(passport) {
		err = ecode.ServerErr
		return
	}
	return s.dao.PassportGet(passport)
}
