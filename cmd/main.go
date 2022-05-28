package main

import (
	"fmt"
	"os"
	"shippo-server/internal/server/http"
	"shippo-server/utils/config"
)

func main() {
	dir, _ := os.Getwd()
	fmt.Printf("main:%v\n", dir)
	config.Init()
	http.New()
}
