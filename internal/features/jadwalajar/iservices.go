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
	var response JadwalajarResponse

	query := `
		SELECT 
			j.id,
			g.nama AS guru,
			m.nama AS mapel,
			k.nama AS kelas,
			j.hari,
			j.jam_mulai,
			j.jam_selesai,
			j.last_editor
		FROM jadwal_ajar j
		JOIN guru g ON g.id = j.guru_id
		JOIN mapel m ON m.id = j.mapel_id
		JOIN kelas k ON k.id = j.kelas_id
		WHERE j.id = ?
	`

	if err := s.DB.Raw(query, id).Scan(&response).Error; err != nil {
		return nil, err
	}

	return &response, nil
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

