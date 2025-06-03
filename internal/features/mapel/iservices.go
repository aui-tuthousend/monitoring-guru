package mapel

import (
	e "monitoring-guru/entities"

	"gorm.io/gorm"
)

type MapelService struct {
	DB *gorm.DB
}

func (s *MapelService) CreateMapel(mapel *e.Mapel) error {
	return s.DB.Create(mapel).Error
}

func (s *MapelService) UpdateMapel(mapel *e.Mapel) error {
	return s.DB.Save(mapel).Error
}

func (s *MapelService) GetAllMapel() ([]MapelResponse, error) {
	var mapelList []MapelResponse
	if err := s.DB.Find(&mapelList).Error; err != nil {
		return nil, err
	}
	return mapelList, nil
}

func (s *MapelService) GetMapelByID(id string) (*MapelResponse, error) {
	var mapel MapelResponse
	if err := s.DB.Where("id = ?", id).First(&mapel).Error; err != nil {
		return nil, err
	}
	return &mapel, nil
}

func (s *MapelService) DeleteMapel(id string) error {
	var mapel e.Mapel
	if err := s.DB.Where("id = ?", id).First(&mapel).Error; err != nil {
		return err
	}
	return s.DB.Delete(&mapel).Error
}
