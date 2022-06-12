package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"shippo-server/internal/model"
	"shippo-server/utils"
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

func (s *WxService) WXGetToken() (token string, err error) {
	if time.Since(s.wxAccessTokenCreatedAt) > time.Hour {
		err = s.WXRefreshToken()
	}
	token = s.wxAccessToken
	return
}

func (s *WxService) AuthCodeToSession(code string) (r model.AuthCodeToSessionResult, err error) {

	err = utils.HttpGetJSON("https://api.weixin.qq.com/sns/jscode2session?appid="+config.Common.MiniProgramAppID+
		"&secret="+config.Common.MiniProgramAppSecret+"&js_code="+code+"&grant_type=authorization_code", &r)

	if err != nil {
		return
	}
	fmt.Printf("AuthCodeToSession:%+v\n", r)

	if r.Errcode != 0 {
		err = errors.New("AuthCodeToSession: " + r.Errmsg)
	}

	return
}
