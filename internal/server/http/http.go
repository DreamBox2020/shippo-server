package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func initServer(address string, router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func Init() {
	engine := gin.Default()
	outerRouter(engine)
	server := initServer(":8233", engine)
	server.ListenAndServe()
}

func outerRouter(Router *gin.Engine) {
	base := Router.Group("")
	initUserRouter(base)

}
