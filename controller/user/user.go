package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login 用户登录
func Login(context *gin.Context) {
	context.JSON(http.StatusOK, "")
}
