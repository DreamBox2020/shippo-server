package configs

type Common struct {
	AppSecret            string `json:"appSecret"`
	AppID                string `json:"appID"`
	WxToken              string `json:"wxToken"`
	MiniProgramAppSecret string `json:"miniProgramAppSecret"`
	MiniProgramAppID     string `json:"miniProgramAppID"`
}
