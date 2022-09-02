package http

import (
	"net/http"
	"shippo-server/internal/service"
	"shippo-server/middleware"
	"shippo-server/utils/box"
	"shippo-server/utils/config"
	"time"
)

type ServerGroup struct {
	User             *UserServer
	Temp             *TempServer
	Passport         *PassportServer
	File             *FileServer
	Captcha          *CaptchaServer
	Role             *RoleServer
	PermissionAccess *PermissionAccessServer
	PermissionPolicy *PermissionPolicyServer
	Wx               *WxServer
	WxArticle        *WxArticleServer
	WxCommentLike    *WxCommentLikeServer
	WxOffiaccount    *WxOffiaccountServer
	WxComment        *WxCommentServer
	WxPassport       *WxPassportServer
	Express          *ExpressServer
}

type Server struct {
	engine  *box.Engine
	router  *box.RouterGroup
	service *service.ServiceGroup
	Group   *ServerGroup
}

func New() *Server {
	var engine = box.New()
	engine.GinEngine.Use(middleware.Cors())
	//engine.MaxMultipartMemory = 8 << 20 // 8 MiB

	var svc = service.New()
	s := &Server{
		engine:  engine,
		router:  engine.Group("/v1"),
		service: svc.Group,
		Group:   nil,
	}
	s.initGroup()
	s.init()

	return s
}

func (t *Server) initGroup() {
	t.Group = &ServerGroup{
		User:             NewUserServer(t),
		Temp:             NewTempServer(t),
		Passport:         NewPassportServer(t),
		File:             NewFileServer(t),
		Captcha:          NewCaptchaServer(t),
		Role:             NewRoleServer(t),
		PermissionAccess: NewPermissionAccessServer(t),
		PermissionPolicy: NewPermissionPolicyServer(t),
		Wx:               NewWxServer(t),
		WxArticle:        NewWxArticleServer(t),
		WxCommentLike:    NewWxCommentLikeServer(t),
		WxOffiaccount:    NewWxOffiaccountServer(t),
		WxComment:        NewWxCommentServer(t),
		WxPassport:       NewWxPassportServer(t),
		Express:          NewExpressServer(t),
	}
}

func (t *Server) init() {

	// 初始化用户信息的中间件
	t.engine.Use(t.Group.Passport.PassportGet)
	t.engine.Use(t.Group.Passport.Auth)

	server := t.initServer(config.Server.Addr, t.engine)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func (t *Server) initServer(address string, router *box.Engine) *http.Server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
