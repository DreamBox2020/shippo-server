package box

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type response struct {
	Code     string      `json:"code"`
	Message  string      `json:"message"`
	Success  bool        `json:"success"`
	Session  string      `json:"session"`
	Resource interface{} `json:"resource"`
	Sign     string      `json:"sign"`
	Other    interface{} `json:"other"`
}

type Context struct {
	ctx *gin.Context
}

func New(ctx *gin.Context) *Context {
	return &Context{
		ctx: ctx,
	}
}

func (context *Context) JSON(data interface{}, err error) {
	context.ctx.JSON(http.StatusOK, &response{
		Resource: data,
	})
}

type HandlerFunc func(context *Context)

func Handler(h HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		h(New(ctx))
	}
}
