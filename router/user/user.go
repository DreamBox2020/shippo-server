package user

import (
	"github.com/gin-gonic/gin"
	"shippo-server/controller/user"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("login", user.Login)
	}
}
