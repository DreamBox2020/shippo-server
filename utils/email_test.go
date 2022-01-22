package utils

import (
	"testing"
)

func TestSendEmail(t *testing.T) {
	SendEmail(emailConf.TestEmail, GenerateCaptcha())
}
