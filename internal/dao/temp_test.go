package dao

import (
	"fmt"
	"testing"
)

func TestTempDao_Temp_trade_20220108_findSuccess(t *testing.T) {
	d := New()
	res, _ := d.Group.Temp.Temp_trade_20220108_findSuccess()
	for _, v := range res {
		fmt.Printf("TestTempDao_Temp_trade_20220108_findSuccess:%+v\n", v)
	}
}
