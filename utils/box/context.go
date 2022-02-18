package box

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"runtime"
	"shippo-server/internal/model"
	"shippo-server/utils/ecode"
	"strings"
)

const (
	AccessAll     = 0 // 无需权限
	AccessLoginOK = 1 // 必须已经登录
	AccessNoLogin = 2 // 必须没有登录
)

var (
	handlers []HandlerFunc
)

type response struct {
	Code     int         `json:"code"`
	Message  string      `json:"message"`
	Success  bool        `json:"success"`
	Session  string      `json:"session"`
	Resource string      `json:"resource"`
	Sign     string      `json:"sign"`
	Other    interface{} `json:"other"`
}

type request struct {
	Passport string      `json:"passport"`
	Session  string      `json:"session"`
	Resource string      `json:"resource"`
	Sign     string      `json:"sign"`
	Other    interface{} `json:"other"`
}

type Context struct {
	index    int
	Ctx      *gin.Context
	Req      *request
	Passport *model.Passport
	Access   int
	IsEnd    bool
}

func Use(middleware ...HandlerFunc) {
	handlers = append(handlers, middleware...)
}

func New(ctx *gin.Context, access int) (bctx *Context) {
	var req *request
	if ctx.GetHeader("Content-Type") == "application/json" {
		ctx.ShouldBindJSON(&req)
	} else {
		req = &request{}
	}
	bctx = &Context{
		index:  -1,
		Ctx:    ctx,
		Req:    req,
		Access: access,
		IsEnd:  false,
	}
	bctx.Next()
	return
}

func (context *Context) Next() {
	context.index++
	for context.index < len(handlers) {
		handlers[context.index](context)
		context.index++
	}
}

func (context *Context) JSON(data interface{}, err error) {
	context.IsEnd = true
	code := ecode.Cause(err)
	res, err2 := json.Marshal(data)
	if err2 != nil {
		fmt.Printf("box->context->JSON->data:%+v\n", data)
		fmt.Printf("box->context->JSON->err2:%+v\n", err2)
	}
	context.Ctx.JSON(http.StatusOK, &response{
		Code:     code.Code(),
		Message:  code.Message(),
		Success:  err == nil,
		Session:  context.Req.Session,
		Resource: string(res),
	})
}

func (context *Context) ShouldBindJSON(obj interface{}) {
	json.Unmarshal([]byte(context.Req.Resource), obj)
}

func (context *Context) Data(contentType string, data []byte, fileName string) {
	context.IsEnd = true
	context.Ctx.Header("content-disposition", `attachment; filename=`+fileName)
	context.Ctx.Data(http.StatusOK, contentType, data)
}

type HandlerFunc func(*Context)

func Handler(h HandlerFunc, access int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bctx := New(ctx, access)
		if !bctx.IsEnd {
			h(bctx)
		}
	}
}

// 获取函数名称
func GetFunctionName(i interface{}, seps ...rune) string {
	fn := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	//fmt.Println(fn)
	fn = strings.TrimSuffix(fn, "-fm")

	// 用 seps 进行分割
	fields := strings.FieldsFunc(fn, func(sep rune) bool {
		for _, s := range seps {
			if sep == s {
				return true
			}
		}
		return false
	})
	//fmt.Println(fields)
	if size := len(fields); size > 0 {
		return fields[size-1]
	}
	return ""
}

type BoxHandler struct {
	obj interface{}
}

func NewBoxHandler(obj interface{}) *BoxHandler {
	return &BoxHandler{obj}
}

func (bh *BoxHandler) H(h HandlerFunc, access int) gin.HandlerFunc {
	funName := GetFunctionName(h, '/', '.')
	//fmt.Printf("GetFunctionName:%+v\n", funName)
	mtV := reflect.ValueOf(bh.obj).Elem()
	m := mtV.MethodByName(funName)

	return Handler(func(bhf *Context) {

		params := make([]reflect.Value, 1)
		params[0] = reflect.ValueOf(bhf)
		m.Call(params)

	}, access)
}
