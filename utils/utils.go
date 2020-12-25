package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadConfigFromFile(path string, config interface{}) error {
	file, _ := os.Open(path)
	defer file.Close()
	bytes, _ := ioutil.ReadAll(file)
	return json.Unmarshal(bytes, &config)
}
