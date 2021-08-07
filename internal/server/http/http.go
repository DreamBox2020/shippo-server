package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shippo-server/configs"
	"shippo-server/internal/service"
	"shippo-server/utils"
	"shippo-server/utils/ecode"
	"time"
)

var (
	svc *service.Service
)

func Init(s *service.Service) {
	svc = s

	var conf configs.Server
	utils.ReadConfigFromFile("configs/server.json", &conf)

	ecode.Register(ecode.Messages)

	engine := gin.Default()
	engine.Use(cors())
	outerRouter(engine)
	server := initServer(conf.Addr, engine)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func outerRouter(Router *gin.Engine) {
	base := Router.Group("")
	initUserRouter(base)
	initFileRouter(base)
	initPassportRouter(base)
	initSmsRouter(base)

}

func cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		} else {
			ctx.Next()
		}
	}
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
