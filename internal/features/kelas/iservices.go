package kelas

import (
	"encoding/json"
	e "monitoring-guru/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type KelasService struct {
	DB *gorm.DB
}

func (s *KelasService) CreateKelas(kelas *e.Kelas, statusKelas *e.StatusKelas) error {
	err := s.DB.Create(kelas).Error
	err2 := s.DB.Create(statusKelas)
	if err != nil || err2 != nil {
		return err
	}
	return nil
}

func (s *KelasService) UpdateKelas(kelas *e.Kelas) error {
	return s.DB.Save(kelas).Error
}

func (s *KelasService) GetAllKelas() ([]KelasResponse, error) {
	var jsonData *string
	query := `
		SELECT json_agg(
			json_build_object(
				'id', k.id,
				'name', k.name,
				'is_active', k.is_active,
				'jurusan', json_build_object(
					'id', j.id,
					'name', j.name
				),
				'ketua_kelas', json_build_object(
					'id', kk.id,
					'nisn', kk.nisn,
					'name', kk.name
				)
			)
		)
		FROM kelas k
		JOIN jurusans j ON k.jurusan_id = j.id::uuid
		JOIN ketua_kelas kk ON k.ketua_id = kk.id::uuid
	`
	if err := s.DB.Raw(query).Scan(&jsonData).Error; err != nil {
		return nil, err
	}
	kelasList := []KelasResponse{}
	if jsonData == nil {
		return kelasList, nil
	}
	if err := json.Unmarshal([]byte(*jsonData), &kelasList); err != nil {
		return nil, err
	}
	return kelasList, nil
}


func (s *KelasService) GetKelasByID(id uuid.UUID) (*KelasResponse, error) {
	var jsonData string
	query:= `
		SELECT json_build_object(
			'id', k.id,
			'name', k.name,
			'is_active', k.is_active,
			'jurusan', json_build_object(
				'id', j.id,
				'name', j.name
			),
			'ketua_kelas', json_build_object(
				'id', kk.id,
				'nisn', kk.nisn,
				'name', kk.name
			)
		)
		FROM kelas k
		JOIN jurusans j ON k.jurusan_id = j.id::uuid
		JOIN ketua_kelas kk ON k.ketua_id = kk.id::uuid
		WHERE k.id = ?::uuid
	`
	if err := s.DB.Raw(query, id).Scan(&jsonData).Error; err != nil {
		return nil, err
	}

	var response KelasResponse
	if err := json.Unmarshal([]byte(jsonData), &response); err != nil {
		return nil, err
	}

	return &response, nil
}


func (s *KelasService) GetKelasByJurusan(jurusanID uuid.UUID) ([]KelasResponse, error) {
	var jsonData *string
	query := `
		SELECT json_agg(
			json_build_object(
				'id', k.id,
				'name', k.name,
				'is_active', k.is_active,
				'jurusan', json_build_object(
					'id', j.id,
					'name', j.name
				),
				'ketua_kelas', json_build_object(
					'id', kk.id,
					'nisn', kk.nisn,
					'name', kk.name
				)
			)
		)
		FROM kelas k
		JOIN jurusans j ON k.jurusan_id = j.id::uuid
		JOIN ketua_kelas kk ON k.ketua_id = kk.id::uuid
		WHERE j.id = ?::uuid
	`

	if err := s.DB.Raw(query).Scan(&jsonData).Error; err != nil {
		return nil, err
	}
	kelasList := []KelasResponse{}
	if jsonData == nil {
		return kelasList, nil
	}
	if err := json.Unmarshal([]byte(*jsonData), &kelasList); err != nil {
		return nil, err
	}
	return kelasList, nil
}

func (s *KelasService) GetKelasByKetuaOrWakil(ketuaKelasID string) (*e.Kelas, error) {
	var kelas e.Kelas
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
