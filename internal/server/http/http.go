package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shippo-server/internal/service"
	"time"
)

var (
	svc *service.Service
)

func Init(s *service.Service) {
	svc = s

	engine := gin.Default()
	outerRouter(engine)
	server := initServer(":8233", engine)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func outerRouter(Router *gin.Engine) {
	base := Router.Group("")
	initUserRouter(base)
}

func initServer(address string, router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
