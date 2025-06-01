package jadwalajar

import (
	e "monitoring-guru/entities"
	"gorm.io/gorm"
)


type JadwalajarService struct {
	DB *gorm.DB
}

func (s *JadwalajarService) CreateJadwalajar(jadwalajar *e.JadwalAjar) error {
	return s.DB.Create(jadwalajar).Error
}

func (s *JadwalajarService) UpdateJadwalajar(jadwalajar *e.JadwalAjar) error {
	return s.DB.Save(jadwalajar).Error
}

func (s *JadwalajarService) GetJadwalajarByID(id string) (*JadwalajarResponse, error) {
	var jadwalajar JadwalajarResponse
	if err := s.DB.Where("id = ?", id).First(&jadwalajar).Error; err != nil {
		return nil, err
	}
	return &jadwalajar, nil
}

func (s *JadwalajarService) GetAllJadwalajar() ([]JadwalajarResponse, error) {
	var jadwalajarList []JadwalajarResponse
	if err := s.DB.
		Preload("Mapel").
		Preload("Guru").
		// Preload("Kelas").
		Find(&jadwalajarList).Error; err != nil {
		return nil, err
	}
	return jadwalajarList, nil
}


func (s *JadwalajarService) GetJadwalajarByIDGuru(id string, hari string) ([]JadwalajarResponse, error) {
	var jadwalajarList []JadwalajarResponse

	if err := s.DB.
		Preload("Mapel").
		Preload("Guru").
		Preload("Kelas").
		Where("guru_id = ? AND hari = ?", id, hari).
		Find(&jadwalajarList).Error; err != nil {
		return nil, err
	}

	return jadwalajarList, nil
}


func (s *JadwalajarService) GetJadwalajarByIDKelas(id string, hari string) ([]JadwalajarResponse, error) {
	var jadwalajarList []JadwalajarResponse

	if err := s.DB.
		Preload("Mapel").
		Preload("Guru").
		// Preload("Kelas").
		Where("kelas_id = ? AND hari = ?", id, hari).
		Find(&jadwalajarList).Error; err != nil {
		return nil, err
	}

	return jadwalajarList, nil
}

