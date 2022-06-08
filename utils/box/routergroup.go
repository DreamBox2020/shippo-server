package box

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IRouter interface {
	IRoutes
	Group(string) *RouterGroup
}

type IRoutes interface {
	Handle(string, string, ...HandlerFunc)
	Any(string, ...HandlerFunc)
	GET(string, ...HandlerFunc)
	POST(string, ...HandlerFunc)
	DELETE(string, ...HandlerFunc)
	PATCH(string, ...HandlerFunc)
	PUT(string, ...HandlerFunc)
	OPTIONS(string, ...HandlerFunc)
	HEAD(string, ...HandlerFunc)

	StaticFile(string, string)
	StaticFileFS(string, string, http.FileSystem)
	Static(string, string)
	StaticFS(string, http.FileSystem)
}

type RouterGroup struct {
	GinGroup *gin.RouterGroup
	engine   *Engine
	root     bool
}

var _ IRouter = &RouterGroup{}

func (group *RouterGroup) Group(relativePath string) *RouterGroup {
	var ginGroup = group.GinGroup
	if group.root {
		ginGroup = group.engine.GinEngine.Group(relativePath)
	} else {
		ginGroup = ginGroup.Group(relativePath)
	}
	return &RouterGroup{
		GinGroup: ginGroup,
		engine:   group.engine,
	}
}

func (group *RouterGroup) Handle(httpMethod, relativePath string, handlers ...HandlerFunc) {
	group.GinGroup.Handle(httpMethod, relativePath, group.engine.HandlersToGinHandlers(handlers...)...)
}

func (group *RouterGroup) POST(relativePath string, handlers ...HandlerFunc) {
	group.GinGroup.POST(relativePath, group.engine.HandlersToGinHandlers(handlers...)...)
}

func (group *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) {
	group.GinGroup.GET(relativePath, group.engine.HandlersToGinHandlers(handlers...)...)
}

func (group *RouterGroup) DELETE(relativePath string, handlers ...HandlerFunc) {
	group.GinGroup.DELETE(relativePath, group.engine.HandlersToGinHandlers(handlers...)...)
}

func (group *RouterGroup) PATCH(relativePath string, handlers ...HandlerFunc) {
	group.GinGroup.PATCH(relativePath, group.engine.HandlersToGinHandlers(handlers...)...)
}

func (group *RouterGroup) PUT(relativePath string, handlers ...HandlerFunc) {
	group.GinGroup.PUT(relativePath, group.engine.HandlersToGinHandlers(handlers...)...)
}

func (group *RouterGroup) OPTIONS(relativePath string, handlers ...HandlerFunc) {
	group.GinGroup.OPTIONS(relativePath, group.engine.HandlersToGinHandlers(handlers...)...)
}

func (group *RouterGroup) HEAD(relativePath string, handlers ...HandlerFunc) {
	group.GinGroup.HEAD(relativePath, group.engine.HandlersToGinHandlers(handlers...)...)
}

func (group *RouterGroup) Any(relativePath string, handlers ...HandlerFunc) {
	group.GinGroup.Any(relativePath, group.engine.HandlersToGinHandlers(handlers...)...)
}

func (group *RouterGroup) StaticFile(relativePath, filepath string) {
	group.GinGroup.StaticFile(relativePath, filepath)
}

func (group *RouterGroup) StaticFileFS(relativePath, filepath string, fs http.FileSystem) {
	group.GinGroup.StaticFileFS(relativePath, filepath, fs)
}

func (group *RouterGroup) Static(relativePath, root string) {
	group.GinGroup.Static(relativePath, root)
}

func (group *RouterGroup) StaticFS(relativePath string, fs http.FileSystem) {
	group.GinGroup.StaticFS(relativePath, fs)
}
