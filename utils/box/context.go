package box

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"mime/multipart"
	"net/http"
	"shippo-server/internal/model"
	"shippo-server/utils/ecode"
)

const abortIndex int8 = math.MaxInt8 / 2

type Response struct {
	Code     int         `json:"code"`
	Message  string      `json:"message"`
	Success  bool        `json:"success"`
	Session  string      `json:"session"`
	Resource string      `json:"resource"`
	Sign     string      `json:"sign"`
	Other    interface{} `json:"other"`
}

type Request struct {
	Passport string      `json:"passport"`
	Session  string      `json:"session"`
	Resource string      `json:"resource"`
	Sign     string      `json:"sign"`
	Other    interface{} `json:"other"`
}

type Context struct {
	Request *http.Request

	index    int8
	engine   *Engine
	Ctx      *gin.Context
	Req      *Request
	Passport *model.Passport
	User     *model.User
}

func (c *Context) Next() {
	handlers := c.engine.handlers
	c.index++
	for c.index < int8(len(handlers)) {
		handlers[c.index](c)
		c.index++
	}
}

func (c *Context) IsAborted() bool {
	return c.index >= abortIndex
}

func (c *Context) Abort() {
	c.index = abortIndex
}

// JSON 响应json格式的数据
func (c *Context) JSON(data interface{}, err error) {
	res, err1 := json.Marshal(data)
	if err1 != nil {
		fmt.Printf("box->context->JSON->data:%+v\n", data)
		fmt.Printf("box->context->JSON->err:%+v\n", err)
		err = err1
	}
	code := ecode.Cause(err)
	c.Ctx.JSON(http.StatusOK, &Response{
		Code:     code.Code(),
		Message:  code.Message(),
		Success:  err == nil,
		Session:  c.Req.Session,
		Resource: string(res),
	})
}

// ShouldBindJSON 解析json格式的数据
func (c *Context) ShouldBindJSON(obj interface{}) error {
	err := json.Unmarshal([]byte(c.Req.Resource), obj)
	if err != nil {
		fmt.Printf("ShouldBindJSON->err:%+v\n", err)
		c.JSON(nil, ecode.ServerErr)
	}
	return err
}

// Data 响应文件格式的数据
func (c *Context) Data(contentType string, data []byte) {
	c.Ctx.Data(http.StatusOK, contentType, data)
}

// DataDownload 响应文件格式的数据，浏览器会直接下载
func (c *Context) DataDownload(contentType string, data []byte, fileName string) {
	c.Ctx.Header("content-disposition", `attachment; filename=`+fileName)
	c.Ctx.Data(http.StatusOK, contentType, data)
}

// NotFound 响应404
func (c *Context) NotFound() {
	c.Ctx.String(http.StatusNotFound, "404 page not found")
}

func (c *Context) ClientIP() string {
	return c.Ctx.ClientIP()
}

func (c *Context) ContentType() string {
	return c.Ctx.ContentType()
}

func (c *Context) Status(code int) {
	c.Ctx.Status(code)
}

func (c *Context) Header(key, value string) {
	c.Ctx.Header(key, value)
}

func (c *Context) GetHeader(key string) string {
	return c.Ctx.GetHeader(key)
}

func (c *Context) SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool) {
	c.Ctx.SetCookie(name, value, maxAge, path, domain, secure, httpOnly)
}

func (c *Context) Cookie(name string) (string, error) {
	return c.Ctx.Cookie(name)
}

func (c *Context) Param(key string) string {
	return c.Ctx.Param(key)
}

func (c *Context) Query(key string) (value string) {
	return c.Ctx.Query(key)
}

func (c *Context) GetQuery(key string) (string, bool) {
	return c.Ctx.GetQuery(key)
}

func (c *Context) PostForm(key string) (value string) {
	return c.Ctx.PostForm(key)
}

func (c *Context) GetPostForm(key string) (string, bool) {
	return c.Ctx.GetPostForm(key)
}

func (c *Context) FormFile(name string) (*multipart.FileHeader, error) {
	return c.Ctx.FormFile(name)
}

func (c *Context) MultipartForm() (*multipart.Form, error) {
	return c.Ctx.MultipartForm()
}

func (c *Context) SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	return c.Ctx.SaveUploadedFile(file, dst)
}
