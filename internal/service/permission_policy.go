package service

import (
	"shippo-server/internal/model"
	"shippo-server/utils"
)

type PermissionPolicyService struct {
	*Service
}

func NewPermissionPolicyService(s *Service) *PermissionPolicyService {
	return &PermissionPolicyService{s}
}

// 按照策略ID查询某个策略信息
func (t *PermissionPolicyService) FindByID(id uint) (p model.PermissionPolicy, err error) {
	p, err = t.dao.PermissionPolicy.FindByID(id)
	return
}

// 按照策略名称查询某个策略信息
func (t *PermissionPolicyService) FindByPolicyName(name string) (p model.PermissionPolicy, err error) {
	p, err = t.dao.PermissionPolicy.FindByPolicyName(name)
	return
}

// 查询某个策略所拥有的访问规则
func (t *PermissionPolicyService) FindPermissionAccessByID(id uint) (
	p []model.PermissionAccess, err error) {
	p, err = t.dao.PermissionPolicy.FindPermissionAccessByID(id)
	return
}

// 查询某个策略所拥有的访问规则
func (t *PermissionPolicyService) FindPermissionAccessByPolicyName(name string) (
	p []model.PermissionAccess, err error) {
	p, err = t.dao.PermissionPolicy.FindPermissionAccessByPolicyName(name)
	return
}

// 按照类型查询某个策略所拥有的访问规则
func (t *PermissionPolicyService) FindPermissionAccessByType(id uint, accessType string) (
	p []model.PermissionAccess, err error) {
	p, err = t.dao.PermissionPolicy.FindPermissionAccessByType(id, accessType)
	return
}

// 按照类型查询某个策略所拥有的访问规则
func (t *PermissionPolicyService) FindPermissionAccessByPolicyNameAndType(name string, accessType string) (
	p []model.PermissionAccess, err error) {
	p, err = t.dao.PermissionPolicy.FindPermissionAccessByPolicyNameAndType(name, accessType)
	return
}

func (t *PermissionPolicyService) PermissionPolicyCreate(p model.PermissionPolicy) (err error) {
	_, err = t.dao.PermissionPolicy.PermissionPolicyCreate(p.PolicyName, p.Remark)
	return
}

func (t *PermissionPolicyService) PermissionPolicyDel(p model.PermissionPolicy) (err error) {
	err = t.dao.PermissionPolicy.PermissionPolicyDel(p.ID)
	return
}

func (t *PermissionPolicyService) PermissionPolicyUpdate(p model.PermissionPolicy) (err error) {
	err = t.dao.PermissionPolicy.PermissionPolicyUpdate(p)
	return
}

func (t *PermissionPolicyService) PermissionPolicyFindAllExtStatus(id uint) (
	p []model.PermissionPolicyStatus, err error) {
	p, err = t.dao.PermissionPolicy.PermissionPolicyFindAllExtStatus(id)
	return
}

func (t *PermissionPolicyService) PermissionPolicyFindAll() (p []model.PermissionPolicyCount, err error) {
	p, err = t.dao.PermissionPolicy.PermissionPolicyFindAll()
	return
}

func (t *PermissionPolicyService) PermissionPolicyFind(p model.PermissionPolicy) (
	list model.PermissionPolicyCount, err error) {
	list, err = t.dao.PermissionPolicy.PermissionPolicyFind(p.ID)
	return
}

// 更新权限关联关系
func (t *PermissionPolicyService) PermissionAssociationUpdate(policyId uint, access []uint) (err error) {

	list, err := t.dao.PermissionPolicy.PermissionAssociationFindPolicyIdList(policyId)
	if err != nil {
		return
	}

	// 如果 旧的列表不包含新的id，就创建
	for _, newAccessId := range access {
		if !utils.In(newAccessId, list) {
			_, err = t.dao.PermissionPolicy.PermissionAssociationCreate(policyId, newAccessId)
			if err != nil {
				return
			}
		}
	}

	// 如果 新的列表不包含旧的id，就删除
	for _, oldAccessId := range list {
		if !utils.In(oldAccessId, access) {
			err = t.dao.PermissionPolicy.PermissionAssociationDel(policyId, oldAccessId)
			if err != nil {
				return
			}
		}
	}

	return
}
