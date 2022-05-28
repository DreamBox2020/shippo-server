package email

import (
	"shippo-server/utils"
	"shippo-server/utils/config"
	"testing"
)

func TestSendEmail(t *testing.T) {
	config.Init()
	SendEmail(config.Email.TestEmail, utils.GenerateCaptcha())
}
