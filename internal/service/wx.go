package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"shippo-server/configs"
	"shippo-server/utils"
	"shippo-server/utils/box"
	"time"
)

type WxService struct {
	*Service
}

func NewWxService(s *Service) *WxService {
	return &WxService{s}
}

func (s *WxService) WXRefreshToken(c *box.Context) (err error) {
	var conf configs.Common
	utils.ReadConfigFromFile("configs/common.json", &conf)

	resp, _ := http.Get("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + conf.AppID + "&secret=" + conf.AppSecret)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var res = new(struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
		Errcode     int    `json:"errcode"`
		Errmsg      string `json:"errmsg"`
	})

	json.Unmarshal(body, res)

	fmt.Printf("RefreshToken: %+v\n", res)

	s.wxAccessToken = res.AccessToken
	s.wxAccessTokenCreatedAt = time.Now()

	return
}

func (s *WxService) WXGetToken(c *box.Context) (token string, err error) {
	if time.Since(s.wxAccessTokenCreatedAt) > time.Hour {
		err = s.WXRefreshToken(c)
	}
	token = s.wxAccessToken
	return
}
