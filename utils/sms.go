package utils

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"shippo-server/configs"
)

var smsConf configs.Sms

func SendSms(phone string, code string) {

	if emailConf.Address == "" {
		if err := ReadConfigFromFile("./configs/sms.json", &smsConf); err != nil {
			panic(err)
		}
	}

	client, err := dysmsapi.NewClientWithAccessKey(smsConf.RegionId, smsConf.AccessKeyId, smsConf.AccessKeySecret)

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = phone
	request.SignName = smsConf.SignName
	request.TemplateCode = smsConf.TemplateCode
	request.TemplateParam = "{\"code\":\"" + code + "\"}"

	response, err := client.SendSms(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("SendSms: %v\n", response)
}
