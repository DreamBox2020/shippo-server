package service

import "shippo-server/internal/model"

type PolicyService struct {
	*Service
}

func NewPolicyService(s *Service) *PolicyService {
	return &PolicyService{s}
}

// 按照策略ID查询某个策略信息
func (t *PolicyService) FindByID(id uint) (p model.PermissionPolicy, err error) {
	p, err = t.dao.Policy.FindByID(id)
	return
}

// 按照策略名称查询某个策略信息
func (t *PolicyService) FindByPolicyName(name string) (p model.PermissionPolicy, err error) {
	p, err = t.dao.Policy.FindByPolicyName(name)
	return
}

// 查询某个策略所拥有的访问规则
func (t *PolicyService) FindPermissionAccessByID(id uint) (
	p []model.PermissionAccess, err error) {
	p, err = t.dao.Policy.FindPermissionAccessByID(id)
	return
}

// 查询某个策略所拥有的访问规则
func (t *PolicyService) FindPermissionAccessByPolicyName(name string) (
	p []model.PermissionAccess, err error) {
	p, err = t.dao.Policy.FindPermissionAccessByPolicyName(name)
	return
}

// 按照类型查询某个策略所拥有的访问规则
func (t *PolicyService) FindPermissionAccessByType(id uint, accessType string) (
	p []model.PermissionAccess, err error) {
	p, err = t.dao.Policy.FindPermissionAccessByType(id, accessType)
	return
}

// 按照类型查询某个策略所拥有的访问规则
func (t *PolicyService) FindPermissionAccessByPolicyNameAndType(name string, accessType string) (
	p []model.PermissionAccess, err error) {
	p, err = t.dao.Policy.FindPermissionAccessByPolicyNameAndType(name, accessType)
	return
}
