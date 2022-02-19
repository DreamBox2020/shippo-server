package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shippo-server/configs"
	"shippo-server/internal/service"
	"shippo-server/middleware"
	"shippo-server/utils"
	"shippo-server/utils/box"
	"shippo-server/utils/ecode"
	"time"
)

type ServerGroup struct {
	User      *UserServer
	Temp      *TempServer
	Passport  *PassportServer
	File      *FileServer
	Captcha   *CaptchaServer
	AdminUser *AdminUserServer
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
		User:      NewUserServer(d),
		Temp:      NewTempServer(d),
		Passport:  NewPassportServer(d),
		File:      NewFileServer(d),
		Captcha:   NewCaptchaServer(d),
		AdminUser: NewAdminUserServer(d),
	}
}

func (s *Server) InitRouter(engine *gin.Engine) {
	router := engine.Group("")
	s.Group.User.InitRouter(router)
	s.Group.Temp.InitRouter(router)
	s.Group.Passport.InitRouter(router)
	s.Group.File.InitRouter(router)
	s.Group.Captcha.InitRouter(router)
	s.Group.AdminUser.InitRouter(router)
}

var serverConf configs.Server

func (s *Server) Init() {
	if err := utils.ReadConfigFromFile("configs/server.json", &serverConf); err != nil {
		panic(err)
	}

	// 初始化错误码
	ecode.Register(ecode.Messages)
	// 初始化用户信息的中间件
	box.Use(func(c *box.Context) {
		s.Group.Passport.PassportGet(c)
	})

	engine := gin.Default()
	//engine.MaxMultipartMemory = 8 << 20 // 8 MiB
	engine.Use(middleware.Cors())
	s.InitRouter(engine)

	server := s.InitServer(serverConf.Addr, engine)
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
