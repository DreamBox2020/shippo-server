package utils

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"shippo-server/configs"
)

func SendSms(phone string, code string) {

	var conf configs.Sms
	ReadConfigFromFile("./configs/sms.json", &conf)

	client, err := dysmsapi.NewClientWithAccessKey(conf.RegionId, conf.AccessKeyId, conf.AccessKeySecret)

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = phone
	request.SignName = conf.SignName
	request.TemplateCode = conf.TemplateCode
	request.TemplateParam = "{\"code\":\"" + code + "\"}"

	response, err := client.SendSms(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("SendSms: %v\n", response)
}
