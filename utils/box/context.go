package box

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
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
	code := ecode.Cause(err)
	res, err2 := json.Marshal(data)
	if err2 != nil {
		fmt.Printf("box->context->JSON->data:%+v\n", data)
		fmt.Printf("box->context->JSON->err2:%+v\n", err2)
		c.JSON(nil, ecode.ServerErr)
		return
	}
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
	return json.Unmarshal([]byte(c.Req.Resource), obj)
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
