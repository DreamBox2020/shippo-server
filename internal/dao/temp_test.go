package dao

import (
	"fmt"
	"shippo-server/internal/model"
	"shippo-server/utils/config"
	"testing"
)

func TestTempDao_Temp_trade_20220108_findSuccess(t *testing.T) {
	d := newTest()
	res, _ := d.Group.Temp.Temp_trade_20220108_findSuccess()
	for _, v := range res {
		fmt.Printf("TestTempDao_Temp_trade_20220108_findSuccess:%+v\n", v)
	}
}

func TestTempDao_Temp_express_20220914_findByQQAndPhone(t *testing.T) {
	config.Init()
	d := newTest()
	res, _ := d.Group.Temp.Temp_express_20220914_findByQQAndPhone(&model.Temp_express_20220914{
		QQ:    "12345",
		Phone: config.Sms.TestPhoneNumber,
	})
	fmt.Printf("TestTempDao_Temp_express_20220914_findByQQAndPhone:%+v\n", res)
}
