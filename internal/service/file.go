package service

import (
	"fmt"
	"net/http"
	"os"
	"shippo-server/internal/model"
	"shippo-server/utils"
	"shippo-server/utils/config"
	"shippo-server/utils/ecode"
	"time"
)

type FileService struct {
	*Service
}

func NewFileService(s *Service) *FileService {
	return &FileService{s}
}

// ToLocalUrl 网络图片转为本地链接
func (t *FileService) ToLocalUrl(url string, channel string) (r *model.Picture, err error) {

	bytes, err := utils.HttpGet(url)
	if err != nil {
		return
	}

	mime := http.DetectContentType(bytes)
	fmt.Printf("ToLocalUrl->mime:%+v\n", mime)

	var fileType string
	if mime == "image/jpeg" {
		fileType = "jpeg"
	} else if mime == "image/png" {
		fileType = "png"
	} else if mime == "image/gif" {
		fileType = "gif"
	} else {
		err = ecode.FileTypeUnknown
		return
	}
	fmt.Printf("ToLocalUrl->fileType:%+v\n", fileType)

	fileUuid := utils.GenerateToken()
	fmt.Printf("fileUpload->fileUuid:%+v\n", fileUuid)

	date := time.Now().Format("2006/01/02")
	fmt.Printf("fileUpload->date:%+v\n", date)

	dir := "/pic/" + channel + "/" + date + "/"

	fileName := fileUuid + "." + fileType
	fmt.Printf("FileUpload->fileName:%+v\n", fileName)

	uri := dir + fileName
	fmt.Printf("fileUpload->uri:%+v\n", uri)

	dst := config.Server.UploadDir + uri
	fmt.Printf("fileUpload->dst:%+v\n", dst)

	err = os.MkdirAll(config.Server.UploadDir+dir, os.ModePerm)
	if err != nil {
		return
	}

	err = utils.SaveFile(bytes, dst)

	picture, err := t.Group.Picture.Create(model.Picture{
		Path:    uri,
		Uri:     uri,
		Name:    fileName,
		Mime:    mime,
		Channel: channel,
	})

	r = &picture

	return
}
