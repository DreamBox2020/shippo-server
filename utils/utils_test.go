package utils

import (
	"fmt"
	"shippo-server/configs"
	"testing"
)

func TestReadConfigFromFile(t *testing.T) {
	var conf configs.DB
	ReadConfigFromFile("configs/db.json", &conf)
	fmt.Printf("配置项内容：%v\n", conf)
}
