package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"net/http"
	"shippo-server/utils/box"
)

type GoogleapisServer struct {
	*Server
	router *box.RouterGroup
}

func NewGoogleapisServer(server *Server) *GoogleapisServer {
	var s = &GoogleapisServer{
		Server: server,
		router: server.router.Group("googleapis"),
	}
	s.initRouter()
	return s
}

func (t *GoogleapisServer) initRouter() {
	t.router.GinGroup.GET("authorize", t.Authorize)
}

func (t *GoogleapisServer) Authorize(c *gin.Context) {

	code := c.Query("code")

	oauth2Cfg := t.service.Googleapis.Oauth2Cfg

	if code != "" {

		t.service.Googleapis.GetUserinfo(code)
		c.String(200, "Code:"+code)

	} else {

		if oauth2Cfg.RedirectURL == "" {

			protocol := "http://"
			if c.Request.TLS != nil {
				protocol = "https://"
			}
			oauth2Cfg.RedirectURL = protocol + c.Request.Host + c.Request.URL.Path
		}

		location := oauth2Cfg.AuthCodeURL("state", oauth2.AccessTypeOnline,
			oauth2.SetAuthURLParam("prompt", "select_account"),
		)

		fmt.Printf("GoogleapisServer->Authorize->location: %+v\n", location)

		c.Redirect(http.StatusMovedPermanently, location)
	}

}
