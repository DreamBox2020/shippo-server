package model

type WxOffiaccount struct {
	Model
	Appid     string `json:"appid"`
	AvatarUrl string `json:"avatarUrl"`
	Nickname  string `json:"nickname"`
	Username  string `json:"username"`
}
