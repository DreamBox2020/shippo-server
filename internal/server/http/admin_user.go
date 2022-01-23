package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shippo-server/utils/box"
)

func initAdminUserRouter(Router *gin.RouterGroup) {
	r := Router.Group("admin/user")
	{
		r.POST("create", box.Handler(userCreateEmail, box.AccessLoginOK))
	}
}

func userCreateEmail(c *box.Context) {
	var param = new(struct {
		Email string `json:"email"`
	})
	c.ShouldBindJSON(&param)
	fmt.Printf("userCreateEmail: %+v\n", param)

	_, err := svc.AdminUserCreateEmail(c, param.Email)
	c.JSON(nil, err)
}
