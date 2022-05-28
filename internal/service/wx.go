package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"shippo-server/utils/box"
	"shippo-server/utils/config"
	"time"
)

type WxService struct {
	*Service
}

func NewWxService(s *Service) *WxService {
	return &WxService{s}
}

func (s *WxService) WXRefreshToken() (err error) {

	resp, _ := http.Get("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" +
		config.Common.AppID + "&secret=" + config.Common.AppSecret)
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
		err = s.WXRefreshToken()
	}
	token = s.wxAccessToken
	return
}
