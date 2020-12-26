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
