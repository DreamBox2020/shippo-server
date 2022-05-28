package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"shippo-server/configs"
	"shippo-server/utils"
	"strings"
)

var (
	Common configs.Common
	DB     configs.DB
	Email  configs.Email
	Server configs.Server
	Sms    configs.Sms
)

type Config struct {
	Env string
}

func New() *Config {
	file, _ := os.Open(".env")
	defer file.Close()
	bytes, _ := ioutil.ReadAll(file)
	env := strings.TrimSpace(string(bytes))
	fmt.Printf("Config->New-<env:%v\n", env)

	c := &Config{Env: env}

	c.Load("common", &Common)
	c.Load("db", &DB)
	c.Load("email", &Email)
	c.Load("server", &Server)
	c.Load("sms", &Sms)

	return c
}

func (t *Config) Load(name string, obj interface{}) {

	path := "./configs/" + name + "." + t.Env + ".json"
	if !utils.IsExist(path) {
		path = "./configs/" + name + ".json"
	}

	file, _ := os.Open(path)
	defer file.Close()
	bytes, _ := ioutil.ReadAll(file)
	json.Unmarshal(bytes, &obj)

	fmt.Printf("Config->Load->%v:%+v\n", name, obj)

}
