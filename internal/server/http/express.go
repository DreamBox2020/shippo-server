package http

import (
	"encoding/json"
	"fmt"
	"shippo-server/utils/box"
	"time"
)

type ExpressServer struct {
	*Server
	router *box.RouterGroup
}

func NewExpressServer(server *Server) *ExpressServer {
	var s = &ExpressServer{
		Server: server,
		router: server.router.Group("express"),
	}
	s.initRouter()
	return s
}

func (t *ExpressServer) initRouter() {
	t.router.GET("sf/searchRoutes", t.SfSearchRoutes)
}

func (t *ExpressServer) SfSearchRoutes(c *box.Context) {
	var data = new(struct {
		Language       string   `json:"language"`
		TrackingType   string   `json:"trackingType"`
		TrackingNumber []string `json:"trackingNumber"`
		MethodType     string   `json:"methodType"`
	})

	var trackingNumber = []string{""}

	data.Language = "0"
	data.TrackingType = "1"
	data.TrackingNumber = trackingNumber
	data.MethodType = "1"

	msgDataBytes, _ := json.Marshal(data)
	msgData := string(msgDataBytes)

	timestamp := time.Now().Unix()
	checkWord := ""
	msgDigest := t.service.Express.GenerateMsgDigest(msgData, timestamp, checkWord)
	fmt.Printf("SfSearchRoutes->msgDigest:%+v\n", msgDigest)

	t.service.Express.SfSearchRoutes(msgData, msgDigest, timestamp)
}
