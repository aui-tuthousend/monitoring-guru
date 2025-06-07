package mapel

import (
	"encoding/json"
	e "monitoring-guru/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MapelService struct {
	DB *gorm.DB
}

func (s *MapelService) CreateMapel(mapel *e.Mapel) error {
	return s.DB.Create(mapel).Error
}

func (s *MapelService) UpdateMapel(mapel *e.Mapel) error {
	return s.DB.Save(mapel).Error
}

func (s *MapelService) GetAllMapel() ([]MapelResponse, error) {
	var jsonData *string

	query := `
		SELECT json_agg(
			json_build_object(
				'id', m.id,
				'name', m.name,
				'jurusan', json_build_object(
					'id', j.id,
					'name', j.name
				)
			)
		)
		FROM mapels m
		JOIN jurusans j ON m.jurusan_id = j.id::uuid
	`

	if err := s.DB.Raw(query).Scan(&jsonData).Error; err != nil {
		return nil, err
	}
	mapelList := []MapelResponse{}
	if jsonData == nil {
		return mapelList, nil
	}
	if err := json.Unmarshal([]byte(*jsonData), &mapelList); err != nil {
		return nil, err
	}
	return mapelList, nil
}


func (s *MapelService) GetMapelByID(id uuid.UUID) (*MapelResponse, error) {
	var jsonData string
	query := `
		SELECT json_build_object(
			'id', m.id,
			'name', m.name,
			'jurusan', json_build_object(
				'id', j.id,
				'name', j.name
			)
		)
		FROM mapels m
		JOIN jurusans j ON m.jurusan_id = j.id::uuid
		WHERE m.id = ?::uuid
	`

	if err := s.DB.Raw(query, id).Scan(&jsonData).Error; err != nil {
		return nil, err
	}

	var response MapelResponse
	if err := json.Unmarshal([]byte(jsonData), &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *MapelService) DeleteMapel(id string) error {
	var mapel e.Mapel
	if err := s.DB.Where("id = ?", id).First(&mapel).Error; err != nil {
		return err
	}
	return s.DB.Delete(&mapel).Error
}


