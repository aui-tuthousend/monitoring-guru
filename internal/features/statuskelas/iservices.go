package statuskelas

import (
	"encoding/json"

	"gorm.io/gorm"
)

type StatusKelasService struct {
	DB *gorm.DB
}

func (s *StatusKelasService) GetAllStatusKelas() ([]StatusKelasResponse, error) {
	var jsonData *string

	query := `
		SELECT json_agg(
			json_build_object(
				'id', sk.id,
				'kelas', json_build_object(
					'id', k.id,
					'name', k.name,
					'grade', k.grade,
					'jurusan', j.name
				),
				'is_active', sk.is_active,
				'mapel', sk.mapel,
				'pengajar', sk.pengajar,
				'ruangan', sk.ruangan
			)
		)
		FROM status_kelas sk
		JOIN kelas k ON sk.kelas_id = k.id::uuid
		JOIN jurusans j ON k.jurusan_id = j.id::uuid
	`

	if err := s.DB.Raw(query).Scan(&jsonData).Error; err != nil {
		return nil, err
	}
	mapelList := []StatusKelasResponse{}
	if jsonData == nil {
		return mapelList, nil
	}
	if err := json.Unmarshal([]byte(*jsonData), &mapelList); err != nil {
		return nil, err
	}
	return mapelList, nil
}
