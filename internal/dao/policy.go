package dao

import "shippo-server/internal/model"

type PolicyDao struct {
	*Dao
}

func NewPolicyDao(s *Dao) *PolicyDao {
	return &PolicyDao{s}
}

// 按照策略ID查询某个策略信息
func (t *PolicyDao) FindByID(id uint) (p model.PermissionPolicy, err error) {
	err = t.db.Where("id", id).Limit(1).Find(&p).Error
	return
}

// 按照策略名称查询某个策略信息
func (t *PolicyDao) FindByPolicyName(name string) (p model.PermissionPolicy, err error) {
	err = t.db.Where("policy_name", name).Limit(1).Find(&p).Error
	return
}

// 查询某个策略所拥有的访问规则
func (t *PolicyDao) FindPermissionAccessByID(id uint) (
	p []model.PermissionAccess, err error) {
	subQuery := t.db.Model(&model.PermissionAssociation{}).
		Select("access_id").Where("policy_id", id)
	err = t.db.Where("id IN (?)", subQuery).Find(&p).Error
	return
}

// 查询某个策略所拥有的访问规则
func (t *PolicyDao) FindPermissionAccessByPolicyName(name string) (
	p []model.PermissionAccess, err error) {
	r, err := t.FindByPolicyName(name)
	if err != nil {
		return
	}
	p, err = t.FindPermissionAccessByID(r.ID)
	return
}

// 按照类型查询某个策略所拥有的访问规则
func (t *PolicyDao) FindPermissionAccessByType(id uint, accessType string) (
	p []model.PermissionAccess, err error) {
	subQuery := t.db.Model(&model.PermissionAssociation{}).
		Select("access_id").Where("policy_id", id)
	err = t.db.Where("id IN (?)", subQuery).
		Where("access_type", accessType).Find(&p).Error
	return
}

// 按照类型查询某个策略所拥有的访问规则
func (t *PolicyDao) FindPermissionAccessByPolicyNameAndType(name string, accessType string) (
	p []model.PermissionAccess, err error) {
	r, err := t.FindByPolicyName(name)
	if err != nil {
		return
	}
	p, err = t.FindPermissionAccessByType(r.ID, accessType)
	return
}
