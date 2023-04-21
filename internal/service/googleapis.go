package service

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	googleapis_oauth2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
	"shippo-server/utils/config"
)

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

func (t *GoogleapisService) GetAccessToken(authCode string) (token *oauth2.Token, err error) {

	ctx := context.Background()
	token, err = t.Oauth2Cfg.Exchange(ctx, authCode)
	if err != nil {
		return
	}

	fmt.Printf("GoogleapisService->GetAccessToken:%+v\n", token)

	return
}

func (t *GoogleapisService) GetUserinfo(code string) (userinfo *googleapis_oauth2.Userinfo, err error) {

	ctx := context.Background()
	token, err := t.GetAccessToken(code)
	if err != nil {
		return
	}

	service, err := googleapis_oauth2.NewService(ctx, option.WithTokenSource(t.Oauth2Cfg.TokenSource(ctx, token)))
	if err != nil {
		return
	}

	userinfo, err = service.Userinfo.Get().Do()
	if err != nil {
		return
	}

	fmt.Printf("GoogleapisService->GetUserinfo:%+v\n", userinfo)

	return
}
