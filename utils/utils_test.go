package utils

import (
	"fmt"
	"shippo-server/testdata"
	"testing"
)

func TestReadConfigFromFile(t *testing.T) {
	var conf testdata.Config
	ReadConfigFromFile("testdata/config.json", &conf)
	fmt.Printf("配置项内容：%+v\n", conf)
}

func TestIn(t *testing.T) {
	var arr = []int{2, 4, 6, 8}
	fmt.Printf("TestIn%+v\n", In(4, arr))
	fmt.Printf("TestIn%+v\n", In(5, arr))
}
