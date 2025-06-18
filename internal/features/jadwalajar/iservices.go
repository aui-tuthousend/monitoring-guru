package jadwalajar

import (
	"encoding/json"
	"fmt"
	e "monitoring-guru/entities"

	"github.com/google/uuid"
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
			j.hari,
			'guru', json_build_object(
				'id', g.id,
				'name', g.name
			),
			'mapel', json_build_object(
				'id', m.id,
				'name', m.name
			),
			'kelas', json_build_object(
				'id', k.id,
				'name', k.name
			),
			'ruangan', json_build_object(
				'id', r.id,
				'name', r.name
			),
			j.jam_mulai,
			j.jam_selesai,
			j.last_editor
		FROM jadwal_ajars j
		JOIN gurus g ON g.id = j.guru_id
		JOIN mapels m ON m.id = j.mapel_id
		JOIN kelas k ON k.id = j.kelas_id
		JOIN ruangans r ON r.id = j.ruangan_id
		WHERE j.id = ?
	`

	if err := s.DB.Raw(query, id).Scan(&response).Error; err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *JadwalajarService) GetAllJadwalajar() ([]JadwalajarResponse, error) {
	var jsonData *string
	query := `
		SELECT json_agg(
			json_build_object(
				'id', j.id,
				'hari', j.hari,
				'guru', json_build_object(
					'id', g.id,
					'name', g.name
			),
			'mapel', json_build_object(
				'id', m.id,
				'name', m.name
			),
			'kelas', json_build_object(
				'id', k.id,
				'name', k.name
			),
			'ruangan', json_build_object(
				'id', r.id,
				'name', r.name
			),
			'jam_mulai', j.jam_mulai,
			'jam_selesai', j.jam_selesai,
			'last_editor', j.last_editor
			)
		)
		FROM jadwal_ajars j
		JOIN gurus g ON g.id = j.guru_id::uuid
		JOIN mapels m ON m.id = j.mapel_id::uuid
		JOIN kelas k ON k.id = j.kelas_id::uuid
		JOIN ruangans r ON r.id = j.ruangan_id::uuid
	`
	if err := s.DB.Raw(query).Scan(&jsonData).Error; err != nil {
		return nil, err
	}
	jadwalajarList := []JadwalajarResponse{}
	if jsonData == nil {
		return jadwalajarList, nil
	}
	if err := json.Unmarshal([]byte(*jsonData), &jadwalajarList); err != nil {
		return nil, err
	}
	return jadwalajarList, nil
}

func (s *JadwalajarService) GetJadwalajarByIDGuru(id uuid.UUID, hari string) ([]JadwalajarResponse, error) {
	var jsonData *string

	whereClause := "WHERE j.guru_id = ?::uuid"
	args := []interface{}{id}

	if hari != "" {
		whereClause += " AND j.hari = ?"
		args = append(args, hari)
	}

	query := fmt.Sprintf(`
		SELECT json_agg(result)
		FROM (
			SELECT json_build_object(
				'id', j.id,
				'guru', json_build_object(
					'id', g.id,
					'name', g.name
				),
				'mapel', json_build_object(
					'id', m.id,
					'name', m.name
				),
				'kelas', json_build_object(
					'id', k.id,
					'name', k.name
				),
				'ruangan', json_build_object(
					'id', r.id,
					'name', r.name
				),
				'hari', j.hari,
				'jam_mulai', j.jam_mulai,
				'jam_selesai', j.jam_selesai
			) AS result
			FROM jadwal_ajars j
			JOIN gurus g ON g.id = j.guru_id::uuid
			JOIN mapels m ON m.id = j.mapel_id::uuid
			JOIN kelas k ON k.id = j.kelas_id::uuid
			JOIN ruangans r ON r.id = j.ruangan_id::uuid
			%s
			ORDER BY j.jam_mulai
		) sub;
	`, whereClause)

	if err := s.DB.Raw(query, args...).Scan(&jsonData).Error; err != nil {
		return nil, err
	}

	jadwalajarList := []JadwalajarResponse{}
	if jsonData == nil {
		return jadwalajarList, nil
	}

	if err := json.Unmarshal([]byte(*jsonData), &jadwalajarList); err != nil {
		return nil, err
	}

	return jadwalajarList, nil
}



func (s *JadwalajarService) GetJadwalajarByIDKelas(id uuid.UUID, hari string) ([]JadwalajarResponse, error) {
	var jsonData *string

	whereClause := "WHERE j.kelas_id = ?::uuid"
	args := []interface{}{id}

	if hari != "" {
		whereClause += " AND j.hari = ?"
		args = append(args, hari)
	}

	query := fmt.Sprintf(`
		SELECT json_agg(result)
		FROM (
			SELECT json_build_object(
				'id', j.id,
				'guru', json_build_object(
					'id', g.id,
					'name', g.name
				),
				'mapel', json_build_object(
					'id', m.id,
					'name', m.name
				),
				'kelas', json_build_object(
					'id', k.id,
					'name', k.name
				),
				'ruangan', json_build_object(
					'id', r.id,
					'name', r.name
				),
				'hari', j.hari,
				'jam_mulai', j.jam_mulai,
				'jam_selesai', j.jam_selesai
			) AS result
			FROM jadwal_ajars j
			JOIN gurus g ON g.id = j.guru_id::uuid
			JOIN mapels m ON m.id = j.mapel_id::uuid
			JOIN kelas k ON k.id = j.kelas_id::uuid
			JOIN ruangans r ON r.id = j.ruangan_id::uuid
			%s
			ORDER BY j.jam_mulai
		) sub;
	`, whereClause)

	if err := s.DB.Raw(query, args...).Scan(&jsonData).Error; err != nil {
		return nil, err
	}

	jadwalajarList := []JadwalajarResponse{}
	if jsonData == nil {
		return jadwalajarList, nil
	}

	if err := json.Unmarshal([]byte(*jsonData), &jadwalajarList); err != nil {
		return nil, err
	}

	return jadwalajarList, nil
}

