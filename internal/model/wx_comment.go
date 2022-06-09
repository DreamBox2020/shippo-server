package model

type WxComment struct {
	Model
	Content        string `json:"content"`
	ArticleId      uint   `json:"articleId"`
	WxPassportId   uint   `json:"wxPassportId"`
	LikeNum        uint   `json:"likeNum"`
	IsElected      uint   `json:"isElected"`
	IsTop          uint   `json:"isTop"`
	ReplyCommentId uint   `json:"replyCommentId"`
}

type WxCommentExtReplyList struct {
	WxComment
	ReplyList []WxComment `json:"replyList"`
}
