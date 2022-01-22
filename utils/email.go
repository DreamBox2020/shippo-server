package utils

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
	"shippo-server/configs"
)

var emailConf configs.Email

func SendEmail(to string, code string) {

	if emailConf.Address == "" {
		if err := ReadConfigFromFile("./configs/email.json", &emailConf); err != nil {
			panic(err)
		}
	}

	d := gomail.NewDialer(emailConf.Host, emailConf.Port, emailConf.Username, emailConf.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage(gomail.SetEncoding(gomail.Base64))
	m.SetHeader("From", m.FormatAddress(emailConf.Address, emailConf.AddressName))
	m.SetHeader("To", to)
	m.SetHeader("Subject", "邮箱验证码(请勿回复此邮件)")
	m.SetBody("text/html", "<span style=\"border-bottom: 1px dashed rgb(204, 204, 204);\">"+code+"</span>")

	if err := d.DialAndSend(m); err != nil {
		fmt.Print(err.Error())
	}
}
