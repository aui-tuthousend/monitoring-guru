package ketuakelas

import (
	"encoding/json"
	e "monitoring-guru/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


type KetuaKelasService struct {
	DB *gorm.DB
}

func (s *KetuaKelasService) ResponseKetuaKelasMapper(ketua *e.KetuaKelas) *KetuaKelasResponse {
	return &KetuaKelasResponse{
		ID: ketua.ID.String(),
		Nisn: ketua.Nisn,
		Name: ketua.Name,
	}
}

func (s *KetuaKelasService) CreateKetuaKelas(ketua *e.KetuaKelas) error {
	return s.DB.Create(ketua).Error
}

func (s *KetuaKelasService) GetAllKetuaKelas() ([]e.KetuaKelas, error) {
	var ketuaKelas []e.KetuaKelas
	if err := s.DB.Find(&ketuaKelas).Error; err != nil {
		return nil, err
	}
	return ketuaKelas, nil
}

func (s *KetuaKelasService) GetUnsignedKetuaKelas() ([]KetuaKelasResponse, error) {
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
		WHERE k.kelas_id IS NULL
	`
	if err := s.DB.Raw(query).Scan(&jsonData).Error; err != nil {
		return nil, err
	}
	ketuaKelasList := []KetuaKelasResponse{}
	if jsonData == nil {
		return ketuaKelasList, nil
	}
	if err := json.Unmarshal([]byte(*jsonData), &ketuaKelasList); err != nil {
		return nil, err
	}
	return ketuaKelasList, nil
}

func (s *KetuaKelasService) GetKetuaKelas(id string) (*e.KetuaKelas, error) {
	var ketua e.KetuaKelas
	if err := s.DB.Where("id = ?", id).First(&ketua).Error; err != nil {
		return nil, err
	}
	return &ketua, nil
}


func (s *KetuaKelasService) GetKetuaKelasByNISN(nisn string) (*e.KetuaKelas, error) {
	var ketua e.KetuaKelas
	if err := s.DB.Where("nisn = ?", nisn).First(&ketua).Error; err != nil {
		return nil, err
	}
	return &ketua, nil
}

func (s *KetuaKelasService) UpdateKetuaKelas(ketua *e.KetuaKelas) error {
	return s.DB.Save(ketua).Error
}

func (s *KetuaKelasService) DeleteKetuaKelas(id uuid.UUID) error {
	var ketua e.KetuaKelas
	if err := s.DB.Where("id = ?", id).First(&ketua).Error; err != nil {
		return err
	}
	return s.DB.Delete(&ketua).Error
}