package service

import "shippo-server/internal/model"

type PictureService struct {
	*Service
}

func NewPictureService(s *Service) *PictureService {
	return &PictureService{s}
}

func (t *PictureService) Create(m model.Picture) (data model.Picture, err error) {
	data, err = t.dao.Picture.Create(m)
	return
}

func (t *PictureService) Delete(id uint) (err error) {
	err = t.dao.Picture.Delete(id)
	return
}

func (t *PictureService) FindByUri(uri string) (r model.Picture, err error) {
	r, err = t.dao.Picture.FindByUri(uri)
	return
}
