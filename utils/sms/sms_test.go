package sms

import (
	"shippo-server/utils"
	"shippo-server/utils/config"
	"testing"
)

func TestSendSms(t *testing.T) {
	config.Init()
	SendSms(config.Sms.TestPhoneNumber, utils.GenerateCaptcha())
}
