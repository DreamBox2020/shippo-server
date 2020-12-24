package utils

import (
	"fmt"
	"shippo-server/configs"
	"testing"
)

func TestSendSms(t *testing.T) {
	var configSms configs.Sms
	ReadConfigFromFile("configs/sms.json", &configSms)
	fmt.Printf("configSms: %v \n", configSms)
	SendSms(configSms.TestPhoneNumber, "888888")
}
