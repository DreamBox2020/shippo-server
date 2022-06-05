package http

import (
	"github.com/gin-gonic/gin"
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
	AdminUser        *AdminUserServer
	Role             *RoleServer
	PermissionAccess *PermissionAccessServer
	PermissionPolicy *PermissionPolicyServer
	Wx               *WxServer
	WxArticle        *WxArticle
}

type Server struct {
	service *service.ServiceGroup
	Group   *ServerGroup
}

func New() *Server {

	var svc = service.New()
	s := &Server{
		service: svc.Group,
		Group:   nil,
	}
	s.Group = NewGroup(s)
	s.Init()

	return s
}

func NewGroup(d *Server) *ServerGroup {
	return &ServerGroup{
		User:             NewUserServer(d),
		Temp:             NewTempServer(d),
		Passport:         NewPassportServer(d),
		File:             NewFileServer(d),
		Captcha:          NewCaptchaServer(d),
		AdminUser:        NewAdminUserServer(d),
		Role:             NewRoleServer(d),
		PermissionAccess: NewPermissionAccessServer(d),
		PermissionPolicy: NewPermissionPolicyServer(d),
		Wx:               NewWxServer(d),
		WxArticle:        NewWxArticle(d),
	}
}

func (s *Server) InitRouter(engine *gin.Engine) {
	router := engine.Group("v1")
	s.Group.User.InitRouter(router)
	s.Group.Temp.InitRouter(router)
	s.Group.Passport.InitRouter(router)
	s.Group.File.InitRouter(router)
	s.Group.Captcha.InitRouter(router)
	s.Group.AdminUser.InitRouter(router)
	s.Group.Role.InitRouter(router)
	s.Group.PermissionAccess.InitRouter(router)
	s.Group.PermissionPolicy.InitRouter(router)
	s.Group.Wx.InitRouter(router)
	s.Group.WxArticle.InitRouter(router)
}

func (s *Server) Init() {

	// 初始化用户信息的中间件
	box.Use(s.Group.Passport.PassportGet)
	box.Use(s.Group.Passport.Auth)

	engine := gin.Default()
	//engine.MaxMultipartMemory = 8 << 20 // 8 MiB
	engine.Use(middleware.Cors())
	s.InitRouter(engine)

	server := s.InitServer(config.Server.Addr, engine)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func (s *Server) InitServer(address string, router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
