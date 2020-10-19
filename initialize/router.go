package initialize

import (
	"github.com/gin-gonic/gin"
	"shippo-server/router/user"
)

func Routers() *gin.Engine {
	var Engine = gin.Default()
	RootRouter := Engine.Group("")
	user.InitUserRouter(RootRouter)
	return Engine
}
