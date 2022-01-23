package dao

import (
	"fmt"
	"shippo-server/utils"
	"testing"
)

func TestDaoCaptchaSmsInsert(t *testing.T) {
	d := New()
	s, err := d.CaptchaSmsInsert("12345678900", utils.GenerateToken())
	if err != nil {
		panic(err)
	}
	fmt.Printf("TestDaoCaptchaSmsInsert:%+v\n", s)
}

func TestDaoCaptchaEmailInsert(t *testing.T) {
	d := New()
	s, err := d.CaptchaEmailInsert("123456@qq.com", utils.GenerateToken())
	if err != nil {
		panic(err)
	}
	fmt.Printf("TestDaoCaptchaEmailInsert:%+v\n", s)
}

func TestDaoCaptchaDel(t *testing.T) {
	d := New()
	err := d.CaptchaDel("12345678900")
	if err != nil {
		panic(err)
	}
	fmt.Printf("TestDaoCaptchaDel:%+v\n", "OK")
}

func TestDaoCaptchaByTargetAndCode(t *testing.T) {
	d := New()
	token := utils.GenerateToken()
	s, _ := d.CaptchaSmsInsert("12345678900", token)
	fmt.Printf("TestDaoCaptchaByTargetAndCode:%+v\n%+v\n", s.Target, s.Code)

	s, err := d.CaptchaByTargetAndCode(s.Target, s.Code, token)
	if err != nil {
		panic(err)
	}
	fmt.Printf("TestDaoCaptchaByTargetAndCode:%+v\n", s)
}
