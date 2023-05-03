package model

type GoogleapisToken struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        string `json:"expires_in"`
	RefreshToken     string `json:"refresh_token"`
	Scope            string `json:"scope"`
	TokenType        string `json:"token_type"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type GoogleapisUserinfoError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type GoogleapisUserinfo struct {
	Id            string                  `json:"id"`
	Email         string                  `json:"email"`
	VerifiedEmail bool                    `json:"verified_email"`
	Name          string                  `json:"name"`
	GivenName     string                  `json:"given_name"`
	FamilyName    string                  `json:"family_name"`
	Picture       string                  `json:"picture"`
	Locale        string                  `json:"locale"`
	Error         GoogleapisUserinfoError `json:"error"`
}
