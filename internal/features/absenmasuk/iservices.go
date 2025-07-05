package absenmasuk

import (
	"fmt"
	e "monitoring-guru/entities"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AbsenMasukService struct {
	DB *gorm.DB
}

func (s *AbsenMasukService) CreateAbsenMasuk(absenMasuk *e.AbsenMasuk) (*CreateAbsenMasukResponse, error) {
	today := time.Now().Truncate(24 * time.Hour)

	var existing e.AbsenMasuk

	err := s.DB.Where("jadwal_ajar_id = ? AND tanggal = ?", absenMasuk.JadwalAjarID, today).First(&existing).Error
	if err == nil {
		return nil, fmt.Errorf("absen masuk untuk jadwal ini hari ini sudah ada")
	} else if err != gorm.ErrRecordNotFound {
		return nil, err
	}

	absenMasuk.ID = uuid.New()
	absenMasuk.Tanggal = today

	if err := s.DB.Create(absenMasuk).Error; err != nil {
		return nil, err
	}

	return &CreateAbsenMasukResponse{
		ID:       absenMasuk.KelasID.String(),
		IsActive: true,
	}, nil
}
