package model

type WxPassport struct {
	Model
	UnionId           string `json:"unionId"`
	MiniProgramOpenId string `json:"miniProgramOpenId"`
	OffiaccountOpenId string `json:"offiaccountOpenId"`
	Nickname          string `json:"nickname"`
	AvatarUrl         string `json:"avatarUrl"`
}
