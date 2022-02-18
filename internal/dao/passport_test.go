package dao

import (
	"fmt"
	"shippo-server/internal/model"
	"testing"
)

func TestDaoPassportGet(t *testing.T) {
	d := New()

	// 先创建一个通行证
	p, err := d.Group.Passport.PassportCreate(model.Passport{
		UserId: 123456,
		Ip:     "127.0.0.1",
		Ua:     "",
		Client: 1,
	})
	if err != nil {
		panic(err)
	}

	// 然后查询
	p, err = d.Group.Passport.PassportGet(p.Token)
	if err != nil {
		panic(err)
	}

	fmt.Printf("TestDaoPassportGet:%+v\n", p)
}

func TestDaoPassportCreate(t *testing.T) {
	d := New()

	p, err := d.Group.Passport.PassportCreate(model.Passport{
		UserId: 123456,
		Ip:     "127.0.0.1",
		Ua:     "",
		Client: 1,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("TestDaoPassportCreate:%+v\n", p)
}

func TestDaoPassportDelete(t *testing.T) {
	d := New()

	p, err := d.Group.Passport.PassportCreate(model.Passport{
		UserId: 123456,
		Ip:     "127.0.0.1",
		Ua:     "",
		Client: 1,
	})
	if err != nil {
		panic(err)
	}

	err = d.Group.Passport.PassportDelete(p.UserId, p.Client)
	if err != nil {
		panic(err)
	}

	p, err = d.Group.Passport.PassportGet(p.Token)
	if err != nil {
		panic(err)
	}

	fmt.Printf("TestDaoPassportDelete:%+v\n", p)
}

func TestDaoPassportUpdate(t *testing.T) {
	d := New()

	p, err := d.Group.Passport.PassportCreate(model.Passport{
		UserId: 123456,
		Ip:     "",
		Ua:     "",
		Client: 1,
	})
	if err != nil {
		panic(err)
	}

	p, err = d.Group.Passport.PassportUpdate(p.Token, model.Passport{Ip: "127.0.0.1"})
	if err != nil {
		panic(err)
	}

	fmt.Printf("TestDaoPassportUpdate:%+v\n", p)
}
