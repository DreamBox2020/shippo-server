package box

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
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
	Ctx *gin.Context
	Req *request
}

func New(ctx *gin.Context) *Context {
	var req *request
	ctx.ShouldBindJSON(&req)
	return &Context{
		Ctx: ctx,
		Req: req,
	}
}

func (context *Context) JSON(data interface{}, err error) {
	res, _ := json.Marshal(data)
	context.Ctx.JSON(http.StatusOK, &response{
		Code:     0,
		Message:  "OK",
		Success:  err == nil,
		Session:  context.Req.Session,
		Resource: string(res),
	})
}

func (context *Context) ShouldBindJSON(obj interface{}) {
	json.Unmarshal([]byte(context.Req.Resource), obj)
}

func (context *Context) Data(contentType string, data []byte, fileName string) {
	context.Ctx.Header("content-disposition", `attachment; filename=`+fileName)
	context.Ctx.Data(http.StatusOK, contentType, data)
}

type HandlerFunc func(context *Context)

func Handler(h HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		h(New(ctx))
	}
}
