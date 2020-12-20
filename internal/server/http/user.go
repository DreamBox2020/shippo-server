package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func initUserRouter(Router *gin.RouterGroup) {
	user := Router.Group("user")
	{
		user.GET("login", userLogin)
	}
}

func userLogin(context *gin.Context) {
	context.JSON(http.StatusOK, "")
}
