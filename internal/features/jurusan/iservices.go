package jurusan

import (
	e "monitoring-guru/entities"
	"gorm.io/gorm"
)


type JurusanService struct {
	DB *gorm.DB
}

func (s *JurusanService) CreateJurusan(jurusan *e.Jurusan) error {
	return s.DB.Create(jurusan).Error
}

func (s *JurusanService) UpdateJurusan(jurusan *e.Jurusan) error {
	return s.DB.Save(jurusan).Error
}

func (s *JurusanService) GetJurusanByID(id string) (*e.Jurusan, error) {
	var jurusan e.Jurusan
	if err := s.DB.Where("id = ?", id).First(&jurusan).Error; err != nil {
		return nil, err
	}
	return &jurusan, nil
}

func (s *JurusanService) IsJurusanExist(kodeJurusan string, name string) bool {
	var jurusan e.Jurusan
	if err := s.DB.Where("kode_jurusan = ? AND name = ?", kodeJurusan, name).First(&jurusan).Error; err != nil {
		return false
	}
	return true
}

func (s *JurusanService) GetAllJurusan() ([]e.Jurusan, error) {
	var jurusanList []e.Jurusan
	if err := s.DB.
		Find(&jurusanList).Error; err != nil {
		return nil, err
	}
	return jurusanList, nil
}

func (s *JurusanService) DeleteJurusan(id string) error {
	var jurusan e.Jurusan
	if err := s.DB.Where("id = ?", id).First(&jurusan).Error; err != nil {
		return err
	}
	return s.DB.Delete(&jurusan).Error
}

func (s *JurusanService) ResponseJurusanMapper(jurusan *e.Jurusan) *JurusanResponse {
	return &JurusanResponse{
		ID: jurusan.ID.String(),
		Name:      jurusan.Name,
		KodeJurusan: jurusan.KodeJurusan,
	}
}