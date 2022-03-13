package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"shippo-server/utils"
	"shippo-server/utils/box"
	"shippo-server/utils/ecode"
	"time"
)

type FileServer struct {
	*Server
}

func NewFileServer(s *Server) *FileServer {
	return &FileServer{s}
}

func (t *FileServer) InitRouter(Router *gin.RouterGroup) {
	r := Router.Group("file")
	{
		r.GET("pic/*filePath", box.Handler(t.FileDownload, box.AccessAll))
		r.POST("upload", box.Handler(t.FileUpload, box.AccessLoginOK))
	}
}

func (t *FileServer) FileDownload(c *box.Context) {
	filePath := c.Ctx.Param("filePath")
	fmt.Printf("FileDownload->filePath:%+v\n", filePath)
	if filePath != "" {
		fmt.Printf("FileDownload->filePath:%+v\n", serverConf.UploadDir+"/pic"+filePath)

		if s, err := os.Stat(serverConf.UploadDir + "/pic" + filePath); err == nil {
			fmt.Printf("FileDownload->fileName:%+v\n", s.Name())
			if !s.IsDir() {
				file, _ := os.Open(serverConf.UploadDir + "/pic" + filePath)
				defer file.Close()
				bytes, _ := ioutil.ReadAll(file)
				if len(bytes) > 0 {
					c.Data(http.DetectContentType(bytes), bytes, s.Name())
					return
				}
			}
		}
	}
	c.Ctx.Data(http.StatusNotFound, "image/png", make([]byte, 0))

}

func (t *FileServer) FileUpload(c *box.Context) {
	file, _ := c.Ctx.FormFile("file")
	fmt.Printf("FileUpload->Filename:%+v\n", file.Filename)

	f, _ := file.Open()
	defer f.Close()
	buffer := make([]byte, 512)
	_, err := f.Read(buffer)
	if err != nil {
		c.JSON(nil, ecode.ServerErr)
		return
	}
	mime := http.DetectContentType(buffer)
	fmt.Printf("fileUpload->mime:%+v\n", mime)

	var fileType string
	if mime == "image/jpeg" {
		fileType = "jpeg"
	} else if mime == "image/png" {
		fileType = "png"
	} else if mime == "image/gif" {
		fileType = "gif"
	} else {
		c.JSON(nil, ecode.ServerErr)
		return
	}
	fmt.Printf("fileUpload->fileType:%+v\n", fileType)

	fileUuid := utils.GenerateToken()
	fmt.Printf("fileUpload->fileUuid:%+v\n", fileUuid)

	date := time.Now().Format("2006/01/02")
	fmt.Printf("fileUpload->date:%+v\n", date)

	dir := "pic/temp/" + date + "/"

	uri := dir + fileUuid + "." + fileType
	fmt.Printf("fileUpload->date:%+v\n", uri)

	dst := serverConf.UploadDir + "/" + uri
	fmt.Printf("fileUpload->dst:%+v\n", dst)

	if utils.IsExist(serverConf.UploadDir + "/" + dir) {
		err := os.MkdirAll(serverConf.UploadDir+"/"+dir, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			c.JSON(nil, ecode.ServerErr)
			return
		}
	}

	// 上传文件至指定目录
	c.Ctx.SaveUploadedFile(file, dst)

	c.JSON(nil, nil)
}
