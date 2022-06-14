package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("Cors")
		origin := ctx.GetHeader("origin")
		if origin == "" {
			origin = "*"
		}
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		} else {
			ctx.Next()
		}
	}
}
