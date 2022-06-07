package model

type WxArticle struct {
	Model
	Title         string `json:"title"`
	Url           string `json:"url"`
	Image1        string `json:"image1"`
	Image2        string `json:"image2"`
	CommentSwitch uint   `json:"commentSwitch"`
	OffiaccountId uint   `json:"offiaccountId"`
	WxPassportId  uint   `json:"wxPassportId"`
}
