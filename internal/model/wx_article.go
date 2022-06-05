package model

type WxArticle struct {
	Model
	Title         string `json:"title"`
	Url           string `json:"url"`
	Image1        string `json:"image1"`
	Image2        string `json:"image2"`
	CommentSwitch uint   `json:"comment_switch"`
	OffiaccountId uint   `json:"offiaccount_id"`
	WxPassportId  uint   `json:"wx_passport_id"`
}
