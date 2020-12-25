package main

import (
	"shippo-server/internal/server/http"
	"shippo-server/internal/service"
)

func main() {
	svr := service.New()
	http.Init(svr)
}
