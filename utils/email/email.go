package email

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
	"shippo-server/utils/config"
)

func SendEmail(to string, code string) {
	if config.IsLocal() {
		return
	}

	Send(to, "邮箱验证码(请勿回复此邮件)", "<span style=\"border-bottom: 1px dashed rgb(204, 204, 204);\">"+code+"</span>")

}

func SendWarningEmail(msg string) {
	if config.IsLocal() {
		return
	}

	Send(config.Email.AdminEmail, "服务器告警邮件", "<span>"+msg+"</span>")

}

func Send(to string, subject string, context string) {
	d := gomail.NewDialer(config.Email.Host, config.Email.Port, config.Email.Username, config.Email.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage(gomail.SetEncoding(gomail.Base64))
	m.SetHeader("From", m.FormatAddress(config.Email.Address, config.Email.AddressName))
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", context)

	if err := d.DialAndSend(m); err != nil {
		fmt.Print(err.Error())
	}
}
