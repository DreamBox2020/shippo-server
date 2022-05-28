package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"shippo-server/internal/model"
	"shippo-server/utils"
	"shippo-server/utils/box"
	"shippo-server/utils/config"
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
		r.GET("pic/*filePath", box.Handler(t.FileDownload))
		r.POST("upload", box.Handler(t.FileUpload))
	}
}

func (t *FileServer) FileDownload(c *box.Context) {
	param := c.Ctx.Param("filePath")
	fmt.Printf("FileDownload->filePath:%+v\n", param)
	if param != "" {
		r, err := t.service.Picture.FindByUri("/pic" + param)
		if err == nil && r.ID != 0 {
			fmt.Printf("FileDownload->filePath:%+v\n", config.Server.UploadDir+r.Path)
			file, _ := os.Open(config.Server.UploadDir + r.Path)
			defer file.Close()
			bytes, _ := ioutil.ReadAll(file)
			if len(bytes) > 0 {
				c.Data(r.Mime, bytes)
				return
			}
		}
	}
	c.NotFound()

}

func (t *FileServer) FileUpload(c *box.Context) {
	header, _ := c.Ctx.FormFile("file")
	channel, ok := c.Ctx.GetPostForm("channel")
	if !ok {
		c.JSON(nil, ecode.ServerErr)
		return
	}

	mime := utils.DetectContentType(header)
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

	dir := "/pic/" + channel + "/" + date + "/"

	fileName := fileUuid + "." + fileType
	fmt.Printf("FileUpload->fileName:%+v\n", header.Filename)

	uri := dir + fileName
	fmt.Printf("fileUpload->uri:%+v\n", uri)

	dst := config.Server.UploadDir + uri
	fmt.Printf("fileUpload->dst:%+v\n", dst)

	if err := os.MkdirAll(config.Server.UploadDir+dir, os.ModePerm); err != nil {
		fmt.Println(err)
		c.JSON(nil, ecode.ServerErr)
		return
	}

	// 上传文件至指定目录
	if err := c.Ctx.SaveUploadedFile(header, dst); err != nil {
		fmt.Println(err)
		c.JSON(nil, ecode.ServerErr)
		return
	}

	data, err := t.service.Picture.Create(model.Picture{
		Path:    uri,
		Uri:     uri,
		Name:    fileName,
		Mime:    mime,
		Channel: channel,
	})

	c.JSON(data, err)
}
