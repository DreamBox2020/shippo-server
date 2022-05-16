package http

import (
	"encoding/xml"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"shippo-server/configs"
	"shippo-server/utils"
	"shippo-server/utils/box"
	"sort"
	"strings"
	"time"
)

type WxServer struct {
	*Server
}

func NewWxServer(s *Server) *WxServer {
	return &WxServer{s}
}

func (t *WxServer) InitRouter(Router *gin.RouterGroup) {
	r := Router.Group("wx")
	{
		r.GET("msg", t.Msg)
		r.POST("msg", t.MsgPost)
	}
}

func (t *WxServer) Msg(c *gin.Context) {
	echostr := c.Query("echostr")

	if !makeSignature(c) {
		fmt.Printf("makeSignature error")
		c.String(200, "")
	} else {
		c.String(200, echostr)
	}
}

func (t *WxServer) MsgPost(c *gin.Context) {

	if !makeSignature(c) {
		fmt.Printf("makeSignature error")
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

var conf configs.Common

func makeSignature(c *gin.Context) bool {

	if conf.AppSecret == "" {
		if err := utils.ReadConfigFromFile("configs/common.json", &conf); err != nil {
			panic(err)
		}
	}

	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	signature := c.Query("signature")

	list := []string{conf.WxToken, timestamp, nonce}
	sort.Strings(list)

	fmt.Printf("makeSignature->sort: %+v\n", list)

	signatureGen := utils.SHA1(strings.Join(list, ""))

	fmt.Printf("makeSignature->signature: %+v\n", signature)
	fmt.Printf("makeSignature->signatureGen: %+v\n", signatureGen)

	return signatureGen == signature
}
