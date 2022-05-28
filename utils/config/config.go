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

const (
	ENV_LOCAL = "local"
	ENV_DEV   = "development"
	ENV_PROD  = "production"
)

var (
	Env    string
	Common configs.Common
	DB     configs.DB
	Email  configs.Email
	Server configs.Server
	Sms    configs.Sms
)

func Init() {
	file, _ := os.Open(".env")
	defer file.Close()
	bytes, _ := ioutil.ReadAll(file)
	Env = strings.TrimSpace(string(bytes))
	fmt.Printf("config->Init->env:%v\n", Env)

	Load("common", &Common)
	Load("db", &DB)
	Load("email", &Email)
	Load("server", &Server)
	Load("sms", &Sms)
}

func Load(name string, obj interface{}) {

	path := "./configs/" + name + "." + Env + ".json"
	if !utils.IsExist(path) {
		path = "./configs/" + name + ".json"
	}

	file, _ := os.Open(path)
	defer file.Close()
	bytes, _ := ioutil.ReadAll(file)
	json.Unmarshal(bytes, &obj)

	fmt.Printf("Config->Load->%v:%+v\n", name, obj)

}

func IsLocal() bool {
	return Env == ENV_LOCAL
}

func IsDev() bool {
	return Env == ENV_DEV
}

func IsProd() bool {
	return Env == ENV_PROD
}
