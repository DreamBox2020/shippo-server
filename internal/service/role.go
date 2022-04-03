package service

import (
	"shippo-server/internal/model"
	"shippo-server/utils"
)

type RoleService struct {
	*Service
}

func NewRoleService(s *Service) *RoleService {
	return &RoleService{s}
}

// 增加⻆⾊
func (t *RoleService) RoleCreate(r model.Role) (err error) {
	_, err = t.dao.Role.RoleCreate(r.RoleName, r.Remark)
	return
}

// 删除⻆⾊
func (t *RoleService) RoleDel(param model.Role) (err error) {
	err = t.dao.Role.RoleDelete(param.ID)
	return
}

// 更新⻆⾊名称和备注
func (t *RoleService) RoleUpdate(param model.Role) (err error) {
	err = t.dao.Role.RoleUpdate(param)
	return
}

// 更新角色关联关系
func (t *RoleService) RoleAssociationUpdate(roleId uint, policies []uint) (err error) {

	list, err := t.dao.Role.RoleAssociationFindPolicyIdList(roleId)
	if err != nil {
		return
	}

	// 如果 旧的列表不包含新的id，就创建
	for _, newPolicyId := range policies {
		if !utils.In(newPolicyId, list) {
			_, err = t.dao.Role.RoleAssociationCreate(roleId, newPolicyId)
			if err != nil {
				return
			}
		}
	}

	// 如果 新的列表不包含旧的id，就删除
	for _, oldPolicyId := range list {
		if !utils.In(oldPolicyId, policies) {
			err = t.dao.Role.RoleAssociationDel(roleId, oldPolicyId)
			if err != nil {
				return
			}
		}
	}

	return
}

// 查询全部⻆⾊
func (t *RoleService) RoleFindAll() (p []model.Role, err error) {
	p, err = t.dao.Role.RoleFindAll()
	return
}

// 查询某个角色所拥有的权限策略
func (t *RoleService) RoleFindPermissionPolicy(id uint) (p []model.PermissionPolicy, err error) {
	p, err = t.dao.Role.RoleFindPermissionPolicy(id)
	return
}

// 查询某个角色所拥有的访问规则
func (t *RoleService) RoleFindPermissionAccess(id uint) (p []model.PermissionAccess, err error) {
	p, err = t.dao.Role.RoleFindPermissionAccess(id)
	return
}

// 按照类型查询某个角色所拥有的访问规则
func (t *RoleService) RoleFindPermissionAccessByType(id uint, accessType string) (
	p []model.PermissionAccess, err error) {
	p, err = t.dao.Role.RoleFindPermissionAccessByType(id, accessType)
	return
}
