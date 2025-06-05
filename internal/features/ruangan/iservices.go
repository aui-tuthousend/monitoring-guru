package ruangan

import (
	e "monitoring-guru/entities"

	"gorm.io/gorm"
)

type RuanganService struct {
	DB *gorm.DB
}

func (s *RuanganService) CreateRuangan(ruangan *e.Ruangan) error {
	return s.DB.Create(ruangan).Error
}

func (s *RuanganService) UpdateRuangan(ruangan *e.Ruangan) error {
	return s.DB.Save(ruangan).Error
}

func (s *RuanganService) GetAllRuangan() ([]RuanganResponse, error) {
	var ruanganList []RuanganResponse
	if err := s.DB.Table("ruangans").Find(&ruanganList).Error; err != nil {
		return nil, err
	}
	return ruanganList, nil
}

func (s *RuanganService) GetRuanganByID(id string) (*RuanganResponse, error) {
	var ruangan RuanganResponse
	if err := s.DB.Table("ruangans").Where("id = ?", id).First(&ruangan).Error; err != nil {
		return nil, err
	}
	return &ruangan, nil
}

func (s *RuanganService) DeleteRuangan(id string) error {
	var ruangan e.Ruangan
	if err := s.DB.Where("id = ?", id).First(&ruangan).Error; err != nil {
		return err
	}
	return s.DB.Delete(&ruangan).Error
}
