package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"shippo-server/utils/box"
)

type FileServer struct {
	*Server
}

func NewFileServer(s *Server) *FileServer {
	return &FileServer{s}
}

func (t *FileServer) InitRouter(Router *gin.RouterGroup) {
	var h = box.NewBoxHandler(&t)
	r := Router.Group("file")
	{
		r.GET("d/:id", h.H(t.FileDownload, box.AccessAll))
		r.POST("upload", h.H(t.FileUpload, box.AccessLoginOK))
	}
}

func (t *FileServer) FileDownload(c *box.Context) {
	file, _ := os.Open("testdata/golang.png")
	defer file.Close()
	bytes, _ := ioutil.ReadAll(file)
	c.Data("image/png", bytes, "golang.png")
}

func (t *FileServer) FileUpload(c *box.Context) {
	file, _ := c.Ctx.FormFile("file")
	fmt.Printf("fileUpload:%+v\n", file.Filename)

	// 上传文件至指定目录
	// c.SaveUploadedFile(file, dst)
}
