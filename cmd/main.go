package main

import (
	"fmt"
	"shippo-server/configs"
	"shippo-server/utils"
)

func main() {
	//core.RunServer()
	//dao.New()
	//http.Init()
	var configSms configs.Sms
	utils.ReadConfigFromFile("configs/sms.json", &configSms)
	fmt.Printf("configSms: %v \n", configSms)
	utils.SendSms(configSms.TestPhoneNumber, "888888")
}
