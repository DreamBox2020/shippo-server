package http

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"shippo-server/utils"
	"shippo-server/utils/box"
	"shippo-server/utils/config"
	"sort"
	"strings"
	"time"
)

type WxServer struct {
	*Server
	router *gin.RouterGroup
}

func NewWxServer(server *Server) *WxServer {
	var s = &WxServer{
		Server: server,
		router: server.router.Group("wx"),
	}
	s.initRouter()
	return s
}

func (t *WxServer) initRouter() {
	t.router.GET("authorize", t.Authorize)
	t.router.GET("msg", t.Msg)
	t.router.POST("msg", t.MsgPost)
}

func (t *WxServer) Authorize(c *gin.Context) {

	code := c.Query("code")

	if code != "" {

		resp, _ := http.Get("https://api.weixin.qq.com/sns/oauth2/access_token?appid=" + config.Common.AppID +
			"&secret=" + config.Common.AppSecret + "&code=" + code + "&grant_type=authorization_code")
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		var res = new(struct {
			AccessToken  string `json:"access_token"`
			ExpiresIn    int    `json:"expires_in"`
			RefreshToken string `json:"refresh_token"`
			Openid       string `json:"openid"`
			Scope        string `json:"scope"`
			Errcode      int    `json:"errcode"`
			Errmsg       string `json:"errmsg"`
		})

		json.Unmarshal(body, res)

		fmt.Printf("Authorize->res: %+v\n", res)

		if res.Errmsg != "" {
			c.String(200, "Code:"+code+"\nErrmsg:"+res.Errmsg)
		} else {
			c.String(200, "Code:"+code+"\nOpenid:"+res.Openid)
		}

	} else {
		location := "https://open.weixin.qq.com/connect/oauth2/authorize?appid=" + config.Common.AppID +
			"&redirect_uri=" + "http://" + c.Request.Host + c.Request.URL.Path + "&response_type=code&scope=snsapi_userinfo&state=STATE#wechat_redirect"

		fmt.Printf("Authorize->location: %+v\n", location)

		c.Redirect(http.StatusMovedPermanently, location)
	}

}

func (t *WxServer) Msg(c *gin.Context) {
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	signature := c.Query("signature")
	echostr := c.Query("echostr")

	if !makeSignature(config.Common.WxToken, timestamp, nonce, signature) {
		fmt.Printf("makeSignature error\n")
		c.String(200, "")
	} else {
		c.String(200, echostr)
	}
}

func (t *WxServer) MsgPost(c *gin.Context) {
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	signature := c.Query("signature")

	if !makeSignature(config.Common.WxToken, timestamp, nonce, signature) {
		fmt.Printf("makeSignature error\n")
		c.String(200, "")
	}

	bytes, _ := ioutil.ReadAll(c.Request.Body)

	wxEvent := new(struct {
		ToUserName   string
		FromUserName string
		CreateTime   int64
		MsgType      string
		Event        string
		EventKey     string
		Ticket       string
	})

	xml.Unmarshal(bytes, wxEvent)

	fmt.Printf("MsgPost: %+v\n", wxEvent)

	c.XML(http.StatusOK, box.H{"ToUserName": wxEvent.FromUserName, "FromUserName": wxEvent.ToUserName, "MsgType": "text", "Content": "成功", "CreateTime": time.Now().Unix()})

}

func makeSignature(token, timestamp, nonce, signature string) bool {

	list := []string{token, timestamp, nonce}
	sort.Strings(list)

	fmt.Printf("makeSignature->sort: %+v\n", list)

	signatureGen := utils.SHA1(strings.Join(list, ""))

	fmt.Printf("makeSignature->signature: %+v\n", signature)
	fmt.Printf("makeSignature->signatureGen: %+v\n", signatureGen)

	return signatureGen == signature
}
