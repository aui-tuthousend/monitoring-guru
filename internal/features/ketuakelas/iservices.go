package ketuakelas

import (
	e "monitoring-guru/entities"
	"gorm.io/gorm"
)


type KetuaKelasService struct {
	DB *gorm.DB
}

func (s *KetuaKelasService) CreateKetuaKelas(ketua *e.KetuaKelas) error {
	return s.DB.Create(ketua).Error
}

func (s *KetuaKelasService) GetKetuaKelas(id string) (*e.KetuaKelas, error) {
	var ketua e.KetuaKelas
	if err := s.DB.Where("id = ?", id).First(&ketua).Error; err != nil {
		return nil, err
	}
	return &ketua, nil
}

func (s *KetuaKelasService) GetKetuaKelasByNISN(nip string) (*e.KetuaKelas, error) {
	var ketua e.KetuaKelas
	if err := s.DB.Where("nisn = ?", nip).First(&ketua).Error; err != nil {
		return nil, err
	}
	return &ketua, nil
}

func (s *KetuaKelasService) GetAllKetuaKelas() ([]KetuaKelasResponse, error) {
	var ketuaKelas []KetuaKelasResponse
	if err := s.DB.Find(&ketuaKelas).Error; err != nil {
		return nil, err
	}
	return ketuaKelas, nil
}

func (s *KetuaKelasService) UpdateKetuaKelas(ketua *e.KetuaKelas) error {
	return s.DB.Save(ketua).Error
}

func (s *KetuaKelasService) DeleteKetuaKelas(id string) error {
	var ketua e.KetuaKelas
	if err := s.DB.Where("id = ?", id).First(&ketua).Error; err != nil {
		return err
	}
	return s.DB.Delete(&ketua).Error
}

