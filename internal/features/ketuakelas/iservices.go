package ketuakelas

import (
	"encoding/json"
	e "monitoring-guru/entities"

	"gorm.io/gorm"
)


type KetuaKelasService struct {
	DB *gorm.DB
}

func (s *KetuaKelasService) CreateKetuaKelas(ketua *e.KetuaKelas) error {
	return s.DB.Create(ketua).Error
}

func (s *KetuaKelasService) GetKetuaKelasByID(id string) (*e.KetuaKelas, error) {
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
	var jsonData *string
	query := `
		SELECT json_agg(
			json_build_object(
				'id', k.id,
				'name', k.name,
				'nisn', k.nisn
			)
		)
		FROM ketua_kelas k
	`
	if err := s.DB.Raw(query).Scan(&jsonData).Error; err != nil {
		return nil, err
	}
	mapelList := []KetuaKelasResponse{}
	if jsonData == nil {
		return mapelList, nil
	}
	if err := json.Unmarshal([]byte(*jsonData), &mapelList); err != nil {
		return nil, err
	}
	return mapelList, nil
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

func (s *KetuaKelasService) ResponseKetuaKelasMapper(ketua *e.KetuaKelas) *KetuaKelasResponse {
	return &KetuaKelasResponse{
		ID: ketua.ID,
		Nisn: ketua.Nisn,
		Name: ketua.Name,
	}
}