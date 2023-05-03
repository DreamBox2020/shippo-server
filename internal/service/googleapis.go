package service

import (
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	googleapis_oauth2 "google.golang.org/api/oauth2/v2"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"shippo-server/internal/model"
	"shippo-server/utils"
	"shippo-server/utils/config"
	"strings"
)

const basePath = "https://www.googleapis.com/"

type GoogleapisService struct {
	*Service
	Oauth2Cfg *oauth2.Config
}

func NewGoogleapisService(s *Service) *GoogleapisService {
	return &GoogleapisService{s,
		&oauth2.Config{
			ClientID:     config.Common.GoogleapisClientID,
			ClientSecret: config.Common.GoogleapisClientSecret,
			RedirectURL:  "",
			Scopes: []string{
				googleapis_oauth2.UserinfoEmailScope,
				googleapis_oauth2.UserinfoProfileScope,
				googleapis_oauth2.OpenIDScope,
			},
			Endpoint: google.Endpoint,
		},
	}
}

func (t *GoogleapisService) GetAccessToken(authCode string) (token *model.GoogleapisToken, err error) {

	v := url.Values{
		"client_id":     {config.Common.GoogleapisClientID},
		"client_secret": {config.Common.GoogleapisClientSecret},
		"code":          {authCode},
		"grant_type":    {"authorization_code"},
		"redirect_uri":  {t.Oauth2Cfg.RedirectURL},
	}

	req, err := http.NewRequest("POST", google.Endpoint.TokenURL, strings.NewReader(v.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client, err := utils.NewProxyClient(config.Common.ProxyURL)
	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 1<<20))

	json.Unmarshal(body, &token)

	if err != nil {
		return
	}

	fmt.Printf("GoogleapisService->GetAccessToken:%+v\n", token)

	return
}

func (t *GoogleapisService) GetUserinfo(code string) (userinfo *model.GoogleapisUserinfo, err error) {

	token, err := t.GetAccessToken(code)
	if err != nil {
		return
	}

	req, err := http.NewRequest("GET", basePath+"oauth2/v2/userinfo?access_token="+token.AccessToken, strings.NewReader(""))
	client, err := utils.NewProxyClient(config.Common.ProxyURL)

	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 1<<20))

	json.Unmarshal(body, &userinfo)

	fmt.Printf("GoogleapisService->GetUserinfo:%+v\n", userinfo)

	return
}
