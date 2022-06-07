package model

type WxCommentLike struct {
	Model
	CommentId uint `json:"commentId"`
	WxPassportId uint `json:"wxPassportId"`
}

