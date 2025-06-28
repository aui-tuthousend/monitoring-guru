package izin

import (
	e "monitoring-guru/entities"

	"gorm.io/gorm"
)

type IzinService struct {
	DB *gorm.DB
}

func (s *IzinService) CreateIzin(izin *e.Izin) error {
	return s.DB.Create(izin).Error
}

func (s *IzinService) GetIzin(id string) (*e.Izin, error) {
	var izin e.Izin
	if err := s.DB.Where("id = ?", id).First(&izin).Error; err != nil {
		return nil, err
	}
	return &izin, nil
}

func (s *IzinService) ResponseIzinMapper(izin *e.Izin) *IzinResponse {
	return &IzinResponse{
		ID:           izin.ID.String(),
		GuruID:       izin.GuruID.String(),
		JadwalAjarID: izin.JadwalAjarID.String(),
		TanggalIzin:  izin.TanggalIzin.Format("2006-01-02"), // Format YYYY-MM-DD
		Pesan:        izin.Pesan,
		Approval:     izin.Approval,
	}
}

func (s *IzinService) GetAllIzin() ([]e.Izin, error) {
	var izins []e.Izin
	if err := s.DB.Find(&izins).Error; err != nil {
		return nil, err
	}
	return izins, nil
}

func (s *IzinService) UpdateIzin(izin *e.Izin) error {
	return s.DB.Save(izin).Error
}

func (s *IzinService) DeleteIzin(id string) error {
	var izin e.Izin
	if err := s.DB.Where("id = ?", id).First(&izin).Error; err != nil {
		return err
	}
	return s.DB.Delete(&izin).Error
}
