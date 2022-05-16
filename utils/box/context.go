package box

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"shippo-server/internal/model"
	"shippo-server/utils/ecode"
)

const (
	AccessAll     = 0 // 无需权限
	AccessLoginOK = 1 // 必须已经登录
	AccessNoLogin = 2 // 必须没有登录
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
	Ctx      *gin.Context
	Req      *Request
	Passport *model.Passport
	User     *model.User
	Access   int
}

type HandlerFunc func(*Context)

type HandlersChain []HandlerFunc

var (
	handlers HandlersChain
)

func Use(middleware ...HandlerFunc) {
	handlers = append(handlers, middleware...)
}

func (c *Context) Next() {
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

func (c *Context) ShouldBindJSON(obj interface{}) error {
	return json.Unmarshal([]byte(c.Req.Resource), obj)
}

func (c *Context) Data(contentType string, data []byte, fileName string) {
	c.Ctx.Header("content-disposition", `attachment; filename=`+fileName)
	c.Ctx.Data(http.StatusOK, contentType, data)
}

func New(ctx *gin.Context, access int) (bctx *Context) {
	bctx = &Context{
		index:  -1,
		Ctx:    ctx,
		Req:    nil,
		Access: access,
	}
	if ctx.GetHeader("Content-Type") == "application/json" {
		if err := ctx.ShouldBindJSON(&bctx.Req); err != nil {
			bctx.JSON(nil, ecode.ServerErr)
			return
		}
	} else {
		bctx.Req = &Request{}
		if passport, err := ctx.Cookie("__PASSPORT"); err == nil {
			bctx.Req.Passport = passport
		}
	}

	bctx.Next()
	return
}

func Handler(h HandlerFunc, access int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bctx := New(ctx, access)
		if !bctx.IsAborted() {
			h(bctx)
		}
	}
}

type H map[string]interface{}

func (h H) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name = xml.Name{
		Space: "",
		Local: "xml",
	}
	if err := e.EncodeToken(start); err != nil {
		return err
	}
	for key, value := range h {
		elem := xml.StartElement{
			Name: xml.Name{Space: "", Local: key},
			Attr: []xml.Attr{},
		}
		if err := e.EncodeElement(value, elem); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
