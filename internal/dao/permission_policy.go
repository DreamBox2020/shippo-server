package dao

import "shippo-server/internal/model"

type PermissionPolicyDao struct {
	*Dao
}

func NewPermissionPolicyDao(s *Dao) *PermissionPolicyDao {
	return &PermissionPolicyDao{s}
}

// 按照策略ID查询某个策略信息
func (t *PermissionPolicyDao) FindByID(id uint) (p model.PermissionPolicy, err error) {
	err = t.db.Where("id", id).Limit(1).Find(&p).Error
	return
}

// 按照策略名称查询某个策略信息
func (t *PermissionPolicyDao) FindByPolicyName(name string) (p model.PermissionPolicy, err error) {
	err = t.db.Where("policy_name", name).Limit(1).Find(&p).Error
	return
}

// 查询某个策略所拥有的访问规则
func (t *PermissionPolicyDao) FindPermissionAccessByID(id uint) (
	p []model.PermissionAccess, err error) {
	subQuery := t.db.Model(&model.PermissionAssociation{}).
		Select("access_id").Where("policy_id", id)
	err = t.db.Where("id IN (?)", subQuery).Find(&p).Error
	return
}

// 查询某个策略所拥有的访问规则
func (t *PermissionPolicyDao) FindPermissionAccessByPolicyName(name string) (
	p []model.PermissionAccess, err error) {
	r, err := t.FindByPolicyName(name)
	if err != nil {
		return
	}
	p, err = t.FindPermissionAccessByID(r.ID)
	return
}

// 按照类型查询某个策略所拥有的访问规则
func (t *PermissionPolicyDao) FindPermissionAccessByType(id uint, accessType string) (
	p []model.PermissionAccess, err error) {
	subQuery := t.db.Model(&model.PermissionAssociation{}).
		Select("access_id").Where("policy_id", id)
	err = t.db.Where("id IN (?)", subQuery).
		Where("access_type", accessType).Find(&p).Error
	return
}

// 按照类型查询某个策略所拥有的访问规则
func (t *PermissionPolicyDao) FindPermissionAccessByPolicyNameAndType(name string, accessType string) (
	p []model.PermissionAccess, err error) {
	r, err := t.FindByPolicyName(name)
	if err != nil {
		return
	}
	p, err = t.FindPermissionAccessByType(r.ID, accessType)
	return
}

// 创建一个策略
func (t *PermissionPolicyDao) PermissionPolicyCreate(PolicyName string, Remark string) (p model.PermissionPolicy, err error) {
	p.PolicyName = PolicyName
	p.Remark = Remark
	err = t.db.Create(&p).Error
	return
}

// 删除一个策略
func (t *PermissionPolicyDao) PermissionPolicyDel(id uint) (err error) {
	err = t.db.Delete(&model.PermissionPolicy{}, id).Error
	return
}

// 更新一个策略
func (t *PermissionPolicyDao) PermissionPolicyUpdate(m model.PermissionPolicy) (err error) {
	err = t.db.Model(&model.PermissionPolicy{}).Where("id", m.ID).Select("policy_name", "remark", "updated_at").Updates(&m).Error
	return
}

// 查询全部策略
func (t *PermissionPolicyDao) PermissionPolicyFindAll() (list []model.PermissionPolicyCount, err error) {
	subQuery := t.db.Model(&model.RoleAssociation{}).Select("policy_id", "COUNT(*) AS roleAssociationCount").Group("policy_id")
	err = t.db.Model(&model.PermissionPolicy{}).Select("*").Joins("Left JOIN (?) temp ON temp.policy_id = id", subQuery).Find(&list).Error
	return
}

// 根据id查询某个策略
func (t *PermissionPolicyDao) PermissionPolicyFind(id uint) (p model.PermissionPolicyCount, err error) {
	err = t.db.Model(&model.PermissionPolicy{}).Where("id", id).Limit(1).Scan(&p).Error
	if err != nil {
		return
	}
	err = t.db.Model(&model.RoleAssociation{}).Where("policy_id", id).Count(&p.RoleAssociationCount).Error
	return
}
