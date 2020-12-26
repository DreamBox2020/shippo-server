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
	ctx *gin.Context
	req *request
}

func New(ctx *gin.Context) *Context {
	return &Context{
		ctx: ctx,
	}
}

func (context *Context) JSON(data interface{}, err error) {
	res, _ := json.Marshal(data)
	context.ctx.JSON(http.StatusOK, &response{
		Code:     0,
		Message:  "OK",
		Success:  err == nil,
		Session:  context.req.Session,
		Resource: string(res),
	})
}

func (context *Context) ShouldBindJSON(obj interface{}) {
	context.ctx.ShouldBindJSON(&context.req)
	json.Unmarshal([]byte(context.req.Resource), obj)
}

type HandlerFunc func(context *Context)

func Handler(h HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		h(New(ctx))
	}
}
