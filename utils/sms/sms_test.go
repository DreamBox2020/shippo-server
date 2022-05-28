package sms

import (
	"shippo-server/utils"
	"testing"
)

func TestSendSms(t *testing.T) {
	SendSms(smsConf.TestPhoneNumber, utils.GenerateCaptcha())
}
