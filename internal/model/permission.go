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
	RoleAssociationCount int64 `json:"roleAssociationCount"`
}

type PermissionAccessCount struct {
	PermissionAccess
	PermissionAssociationCount int64 `json:"permissionAssociationCount"`
}
