package sms

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"shippo-server/utils/config"
)

func SendSms(phone string, code string) bool {
	if config.IsLocal() {
		return true
	}

	client, err := dysmsapi.NewClientWithAccessKey(config.Sms.RegionId, config.Sms.AccessKeyId, config.Sms.AccessKeySecret)

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = phone
	request.SignName = config.Sms.SignName
	request.TemplateCode = config.Sms.TemplateCode
	request.TemplateParam = "{\"code\":\"" + code + "\"}"

	response, err := client.SendSms(request)
	if err != nil {
		fmt.Printf("SendSms->err: %v\n", err)
	}
	fmt.Printf("SendSms: %v\n", response)
	return response.Code == "OK"
}
