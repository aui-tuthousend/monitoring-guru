package absenkeluar

import (
	e "monitoring-guru/entities"

	"gorm.io/gorm"
)

type AbsenKeluarService struct {
	DB *gorm.DB
}

func (s *AbsenKeluarService) CreateAbsenKeluar(absenKeluar *e.AbsenKeluar) (error) {
	
	if err := s.DB.Create(absenKeluar).Error; err != nil {
		return err
	}

	return nil
}