package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(context *gin.Context) {
	context.JSON(http.StatusOK, "")
}
