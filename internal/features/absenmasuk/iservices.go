package absenmasuk

import (
	e "monitoring-guru/entities"

	"gorm.io/gorm"
)

type AbsenMasukService struct {
	DB *gorm.DB
}

func (s *AbsenMasukService) CreateAbsenMasuk(absenMasuk *e.AbsenMasuk) (*CreateAbsenMasukResponse, error) {
	
	if err := s.DB.Create(absenMasuk).Error; err != nil {
		return nil, err
	}

	return &CreateAbsenMasukResponse{
		ID: absenMasuk.KelasID.String(),
		IsActive: true,
	}, nil
}
