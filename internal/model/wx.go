package model

type AuthCodeToSessionResult struct {
	Openid     string `json:"openid"`      //用户唯一标识
	SessionKey string `json:"session_key"` //会话密钥
	Unionid    string `json:"unionid"`     //用户在开放平台的唯一标识符，若当前小程序已绑定到微信开放平台帐号下会返回，详见 UnionID 机制说明。
	Errcode    int    `json:"errcode"`     //错误码
	Errmsg     string `json:"errmsg"`      //错误信息
}
