package email

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
	"shippo-server/utils/config"
)

func SendEmail(to string, code string) {

	d := gomail.NewDialer(config.Email.Host, config.Email.Port, config.Email.Username, config.Email.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage(gomail.SetEncoding(gomail.Base64))
	m.SetHeader("From", m.FormatAddress(config.Email.Address, config.Email.AddressName))
	m.SetHeader("To", to)
	m.SetHeader("Subject", "邮箱验证码(请勿回复此邮件)")
	m.SetBody("text/html", "<span style=\"border-bottom: 1px dashed rgb(204, 204, 204);\">"+code+"</span>")

	if err := d.DialAndSend(m); err != nil {
		fmt.Print(err.Error())
	}
}
