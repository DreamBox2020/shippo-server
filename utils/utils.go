package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadConfigFromFile(path string, config interface{}) error {
	file, _ := os.Open(path)
	defer file.Close()
	bytes, _ := ioutil.ReadAll(file)
	fmt.Println("ReadConfigFromFile", bytes)
	return json.Unmarshal(bytes, &config)
}
