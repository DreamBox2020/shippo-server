package model

type WxOffiaccount struct {
	Model
	Appid     int    `json:"appid"`
	AvatarUrl string `json:"avatarUrl"`
	Nickname  string `json:"nickname"`
	Username  string `json:"username"`
}
