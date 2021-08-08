package http

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"shippo-server/utils/box"
)

func initFileRouter(Router *gin.RouterGroup) {
	r := Router.Group("file")
	{
		r.GET("d/:id", box.Handler(fileDownload, box.AccessAll))
	}
}

func fileDownload(c *box.Context) {
	file, _ := os.Open("testdata/golang.png")
	defer file.Close()
	bytes, _ := ioutil.ReadAll(file)
	c.Data("image/png", bytes, "golang.png")
}
