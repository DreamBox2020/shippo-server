package box

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shippo-server/utils/ecode"
)

type HandlerFunc func(*Context)

type HandlersChain []HandlerFunc

type Engine struct {
	RouterGroup
	handlers  HandlersChain
	GinEngine *gin.Engine
}

var _ IRouter = &Engine{}

func New() *Engine {
	engine := &Engine{
		RouterGroup: RouterGroup{
			root: true,
		},
		GinEngine: gin.Default(),
	}
	engine.RouterGroup.engine = engine

	return engine
}

func (engine *Engine) Use(middleware ...HandlerFunc) {
	engine.handlers = append(engine.handlers, middleware...)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	engine.GinEngine.ServeHTTP(w, req)
}

func (engine *Engine) HandlerToGinHandler(handler HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		context := engine.allocateContext(ctx)
		if !context.IsAborted() {
			handler(context)
		}
	}
}

func (engine *Engine) HandlersToGinHandlers(middleware ...HandlerFunc) []gin.HandlerFunc {
	var l = len(middleware)
	var ginHandlers = make([]gin.HandlerFunc, l)
	for i := 0; i < l; i++ {
		ginHandlers[i] = engine.HandlerToGinHandler(middleware[i])
	}
	return ginHandlers
}

func (engine *Engine) allocateContext(ctx *gin.Context) (context *Context) {
	context = &Context{
		Request: ctx.Request,
		index:   -1,
		engine:  engine,
		Ctx:     ctx,
		Req:     nil,
	}
	if ctx.GetHeader("Content-Type") == "application/json" {
		if err := ctx.ShouldBindJSON(&context.Req); err != nil {
			context.JSON(nil, ecode.ServerErr)
			return
		}
	} else {
		context.Req = &Request{}
		if passport, err := ctx.Cookie("__PASSPORT"); err == nil {
			context.Req.Passport = passport
		}
	}

	context.Next()
	return
}
