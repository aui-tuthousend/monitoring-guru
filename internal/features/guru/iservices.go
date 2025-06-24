package guru

import (
	e "monitoring-guru/entities"
	"gorm.io/gorm"
)


type GuruService struct {
	DB *gorm.DB
}

func (s *GuruService) CreateGuru(guru *e.Guru) error {
	return s.DB.Create(guru).Error
}

func (s *GuruService) GetGuru(id string) (*e.Guru, error) {
	var guru e.Guru
	if err := s.DB.Where("id = ?", id).First(&guru).Error; err != nil {
		return nil, err
	}
	return &guru, nil
}

func (s *GuruService) ResponseGuruMapper(guru *e.Guru) *GuruResponse {
	return &GuruResponse{
		ID:      guru.ID.String(),
		Nip:     guru.Nip,
		Name:    guru.Name,
		Jabatan: guru.Jabatan,
	}
}

func (s *GuruService) GetGuruByNIP(nip string) (*e.Guru, error) {
	var guru e.Guru
	if err := s.DB.Where("nip = ?", nip).First(&guru).Error; err != nil {
		return nil, err
	}
	return &guru, nil
}

func (s *GuruService) IsGuruExist(nip string, name string) bool {
	var guru e.Guru
	if err := s.DB.Where("nip = ? AND name = ?", nip, name).First(&guru).Error; err != nil {
		return false
	}
	return true
}

func (s *GuruService) GetAllGuru() ([]e.Guru, error) {
	var gurus []e.Guru
	if err := s.DB.Find(&gurus).Error; err != nil {
		return nil, err
	}
	return gurus, nil
}

func (s *GuruService) UpdateGuru(guru *e.Guru) error {
	return s.DB.Save(guru).Error
}

func (s *GuruService) DeleteGuru(id string) error {
	var guru e.Guru
	if err := s.DB.Where("id = ?", id).First(&guru).Error; err != nil {
		return err
	}
	return s.DB.Delete(&guru).Error
}