package utils

import (
	"testing"
)

func TestSendSms(t *testing.T) {
	SendSms(smsConf.TestPhoneNumber, GenerateCaptcha())
}
