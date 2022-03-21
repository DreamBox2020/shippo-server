package dao

import "shippo-server/internal/model"

type PermissionAccessDao struct {
	*Dao
}

func NewPermissionAccessDao(s *Dao) *PermissionAccessDao {
	return &PermissionAccessDao{s}
}

func (t *PermissionAccessDao) PermissionAccessCreate(p model.PermissionAccess) (err error) {
	var m = &model.PermissionAccess{AccessRule: p.AccessRule, Remark: p.Remark, AccessType: p.AccessType}
	err = t.db.Create(&m).Error
	return
}

func (t *PermissionAccessDao) PermissionAccessDel(id uint) (err error) {
	err = t.db.Delete(&model.PermissionAccess{}, id).Error
	return
}

func (t *PermissionAccessDao) PermissionAccessUpdate(m model.PermissionAccess) (err error) {
	err = t.db.Model(&model.PermissionAccess{}).Where("id", m.ID).Select("access_rule", "remark", "access_type", "updated_at").Updates(&m).Error
	return
}

func (t *PermissionAccessDao) PermissionAccessFindAll() (list []*model.PermissionAccessCount, err error) {
	subQuery := t.db.Model(&model.PermissionAssociation{}).Select("access_id", "COUNT(*) AS PermissionAssociationCount").Group("access_id")
	t.db.Model(&model.PermissionAccess{}).Select("*").Joins("Left JOIN (?) temp ON temp.access_id = id", subQuery).Find(&list)
	return
}

func (t *PermissionAccessDao) PermissionAccessFind(id uint) (p model.PermissionAccessCount, err error) {
	err = t.db.Model(&model.PermissionAccess{}).Where("id", id).Limit(1).Scan(&p).Error
	if err != nil {
		return
	}
	err = t.db.Model(&model.PermissionAssociation{}).Where("access_id", id).Count(&p.PermissionAssociationCount).Error
	return
}
