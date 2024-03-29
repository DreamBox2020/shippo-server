package model

type PermissionPolicy struct {
	Model
	PolicyName string `json:"policyName"`
	Remark     string `json:"remark"`
}

type PermissionAccess struct {
	Model
	AccessRule string `json:"accessRule"`
	Remark     string `json:"remark"`
	AccessType string `json:"accessType"`
}

type PermissionAssociation struct {
	Model
	PolicyId uint
	AccessId uint
}

type PermissionPolicyStatus struct {
	PermissionPolicy
	Status bool `json:"status" gorm:"column:status"`
}

type PermissionAccessStatus struct {
	PermissionAccess
	Status bool `json:"status" gorm:"column:status"`
}

type PermissionPolicyCount struct {
	PermissionPolicy
	RoleAssociationCount int64 `json:"roleAssociationCount" gorm:"column:roleAssociationCount"`
}

type PermissionAccessCount struct {
	PermissionAccess
	PermissionAssociationCount int64 `json:"permissionAssociationCount" gorm:"column:permissionAssociationCount"`
}
