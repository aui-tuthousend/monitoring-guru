package kelas

import (
	e "monitoring-guru/entities"

	"gorm.io/gorm"
)

type KelasService struct {
	DB *gorm.DB
}

func (s *KelasService) CreateKelas(kelas *e.Kelas) error {
	return s.DB.Create(kelas).Error
}

func (s *KelasService) UpdateKelas(kelas *e.Kelas) error {
	return s.DB.Save(kelas).Error
}

func (s *KelasService) GetAllKelas() ([]KelasResponse, error) {
	var kelasList []KelasResponse
	if err := s.DB.Find(&kelasList).Error; err != nil {
		return nil, err
	}
	return kelasList, nil
}

func (s *KelasService) GetKelasByID(id string) (*KelasResponse, error) {
	var kelas KelasResponse
	if err := s.DB.Where("id = ?", id).First(&kelas).Error; err != nil {
		return nil, err
	}
	return &kelas, nil
}

func (s *KelasService) GetKelasByJurusan(jurusanID string) ([]KelasResponse, error) {
	var kelasList []KelasResponse
	if err := s.DB.
		Preload("KetuaKelas").
		Preload("WakilKelas").
		Where("jurusan_id = ?", jurusanID).
		Find(&kelasList).Error; err != nil {
		return nil, err
	}
	return kelasList, nil
}

func (s *KelasService) GetKelasByKetuaOrWakil(ketuaKelasID string) (*KelasResponse, error) {
	var kelas KelasResponse
	if err := s.DB.
		Preload("KetuaKelas").
		Preload("WakilKelas").
		Where("ketua_kelas_id = ? OR wakil_kelas_id = ?", ketuaKelasID, ketuaKelasID).
		First(&kelas).Error; err != nil {
		return nil, err
	}
	return &kelas, nil
}

func (s *KelasService) DeleteKelas(id string) error {
	var kelas e.Kelas
	if err := s.DB.Where("id = ?", id).First(&kelas).Error; err != nil {
		return err
	}
	return s.DB.Delete(&kelas).Error
}
