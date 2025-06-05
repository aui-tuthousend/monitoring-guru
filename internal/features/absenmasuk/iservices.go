package absenmasuk

import (
	e "monitoring-guru/entities"

	"gorm.io/gorm"
)

type AbsenMasukService struct {
	DB *gorm.DB
}

func (s *AbsenMasukService) CreateAbsenMasuk(absenMasuk *e.AbsenMasuk) error {
	return s.DB.Create(absenMasuk).Error
}
