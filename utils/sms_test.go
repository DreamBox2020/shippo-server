package utils

import (
	"shippo-server/configs"
	"testing"
)

func TestSendSms(t *testing.T) {
	var configSms configs.Sms
	ReadConfigFromFile("configs/sms.json", &configSms)
	SendSms(configSms.TestPhoneNumber, "888888")
}
