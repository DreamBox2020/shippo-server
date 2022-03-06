package model

type Role struct {
	Model
	RoleName string `json:"roleName"`
	Remark   string `json:"remark"`
}

type RoleAssociation struct {
	Model
	RoleId   uint
	PolicyId uint
}
