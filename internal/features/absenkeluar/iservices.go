package absenkeluar

import (
	"fmt"
	e "monitoring-guru/entities"
	"time"

	"gorm.io/gorm"
)

type AbsenKeluarService struct {
	DB *gorm.DB
}

func (s *AbsenKeluarService) CreateAbsenKeluar(absenKeluar *e.AbsenKeluar) error {
	today := time.Now().Truncate(24 * time.Hour)

	var existing e.AbsenKeluar

	err := s.DB.Where("absen_masuk_id = ? AND DATE(created_at) = ?", absenKeluar.AbsenMasukID, today).First(&existing).Error

	if err == nil {
		return fmt.Errorf("absen keluar untuk absen masuk ini hari ini sudah ada")
	} else if err != gorm.ErrRecordNotFound {
		return err
	}

	if err := s.DB.Create(absenKeluar).Error; err != nil {
		return err
	}

	return nil
}
