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
	PolicyId int
	AccessId int
}

type PermissionPolicyCount struct {
	PermissionPolicy
	PermissionPolicyCount int64 `json:"permissionAccessCount"`
}

type PermissionAccessCount struct {
	PermissionAccess
	PermissionAccessCount int64 `json:"permissionAccessCount"`
}
