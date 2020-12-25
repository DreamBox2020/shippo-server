package utils

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"shippo-server/configs"
)

func SendSms(phone string, code string) {
	var configSms configs.Sms
	ReadConfigFromFile("./configs/sms.json", &configSms)

	client, err := dysmsapi.NewClientWithAccessKey(configSms.RegionId, configSms.AccessKeyId, configSms.AccessKeySecret)

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = phone
	request.SignName = configSms.SignName
	request.TemplateCode = configSms.TemplateCode
	request.TemplateParam = "{\"code\":\"" + code + "\"}"

	response, err := client.SendSms(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("SendSms: %v\n", response)
}
