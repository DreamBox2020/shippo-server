package dao

import (
	"fmt"
	"shippo-server/internal/model"
	"shippo-server/utils"
	"testing"
)

func TestDaoCaptchaSmsInsert(t *testing.T) {
	d := newTest()
	s, err := d.Group.Captcha.CaptchaSmsInsert("12345678900", utils.GenerateToken(), "")
	if err != nil {
		panic(err)
	}
	fmt.Printf("TestDaoCaptchaSmsInsert:%+v\n", s)
}

func TestDaoCaptchaEmailInsert(t *testing.T) {
	d := newTest()
	s, err := d.Group.Captcha.CaptchaEmailInsert("123456@qq.com", utils.GenerateToken(), "")
	if err != nil {
		panic(err)
	}
	fmt.Printf("TestDaoCaptchaEmailInsert:%+v\n", s)
}

func TestDaoCaptchaDel(t *testing.T) {
	d := newTest()
	err := d.Group.Captcha.CaptchaDel("12345678900")
	if err != nil {
		panic(err)
	}
	fmt.Printf("TestDaoCaptchaDel:%+v\n", "OK")
}

func TestDaoCaptchaByTargetAndCode(t *testing.T) {
	d := newTest()
	token := utils.GenerateToken()
	s, _ := d.Group.Captcha.CaptchaSmsInsert("12345678900", token, "")
	fmt.Printf("TestDaoCaptchaByTargetAndCode:%+v\n%+v\n", s.Target, s.Code)

	s, err := d.Group.Captcha.FindByTargetAndCode(&model.Captcha{
		Target:  s.Target,
		Code:    s.Code,
		Token:   token,
		Channel: "login",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("TestDaoCaptchaByTargetAndCode:%+v\n", s)
}
