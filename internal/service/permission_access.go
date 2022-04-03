package service

import "shippo-server/internal/model"

type PermissionAccessService struct {
	*Service
}

func NewPermissionAccessService(s *Service) *PermissionAccessService {
	return &PermissionAccessService{s}
}

func (t *PermissionAccessService) PermissionAccessCreate(p model.PermissionAccess) (err error) {
	err = t.dao.PermissionAccess.PermissionAccessCreate(p)
	return
}

func (t *PermissionAccessService) PermissionAccessDel(p model.PermissionAccess) (err error) {
	err = t.dao.PermissionAccess.PermissionAccessDel(p.ID)
	return
}

func (t *PermissionAccessService) PermissionAccessUpdate(p model.PermissionAccess) (err error) {
	err = t.dao.PermissionAccess.PermissionAccessUpdate(p)
	return
}

func (t *PermissionAccessService) PermissionAccessFindAllExtStatus(id uint) (
	p []model.PermissionAccessStatus, err error) {
	p, err = t.dao.PermissionAccess.PermissionAccessFindAllExtStatus(id)
	return
}

func (t *PermissionAccessService) PermissionAccessFindAll() (p []model.PermissionAccessCount, err error) {
	p, err = t.dao.PermissionAccess.PermissionAccessFindAll()
	return
}

func (t *PermissionAccessService) PermissionAccessFind(p model.PermissionAccess) (
	list model.PermissionAccessCount, err error) {
	list, err = t.dao.PermissionAccess.PermissionAccessFind(p.ID)
	return
}
