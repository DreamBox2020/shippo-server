package model

type User struct {
	Model
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Exp      uint   `json:"exp"`
	Coin     uint   `json:"coin"`
	Role     uint   `json:"role"`
}

type UserLoginParam struct {
	Phone string `json:"phone"`
	Email string `json:"email"`
	Code  string `json:"code"`
}

type UserFindAllReq struct {
	Pagination
	ID       uint   `json:"id"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
}

type UserFindAllResp struct {
	Pagination
	Items []struct {
		User
		RoleName string `json:"roleName"`
	} `json:"items"`
}
