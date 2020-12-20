package initialize

import (
	"github.com/gin-gonic/gin"
	"shippo-server/router/user"
)

func Routers() *gin.Engine {
	RootRouter := Engine.Group("")
	user.InitUserRouter(RootRouter)
	return Engine
}

var Engine = gin.Default()
