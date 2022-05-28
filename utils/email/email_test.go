package email

import (
	"shippo-server/utils"
	"testing"
)

func TestSendEmail(t *testing.T) {
	SendEmail(emailConf.TestEmail, utils.GenerateCaptcha())
}
