package dao

import (
	"fmt"
	"testing"
)

func TestDaoSmsInsert(t *testing.T) {
	d := New()
	s, err := d.SmsInsert("12345678900")
	if err != nil {
		panic(err)
	}
	fmt.Printf("TestDaoSmsInsert:%+v\n", s)
}

func TestDaoSmsDel(t *testing.T) {
	d := New()
	err := d.SmsDel("12345678900")
	if err != nil {
		panic(err)
	}
	fmt.Printf("TestDaoSmsDel:%+v\n", "OK")
}

func TestDaoSmsByPhoneAndCode(t *testing.T) {
	d := New()
	s, _ := d.SmsInsert("12345678900")
	fmt.Printf("TestDaoSmsByPhoneAndCode:%+v\n%+v\n", s.Phone, s.Code)

	s, err := d.SmsByPhoneAndCode(s.Phone, s.Code)
	if err != nil {
		panic(err)
	}
	fmt.Printf("TestDaoSmsByPhoneAndCode:%+v\n", s)
}
