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
	var jsonData string

	query := `
		SELECT json_build_object(
			'id', j.id,
			'hari', j.hari,
			'guru', json_build_object(
				'id', g.id,
				'name', g.name,
				'nip', g.nip
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
		FROM jadwal_ajars j
		JOIN gurus g ON g.id = j.guru_id
		JOIN mapels m ON m.id = j.mapel_id
		JOIN kelas k ON k.id = j.kelas_id
		JOIN ruangans r ON r.id = j.ruangan_id
		WHERE j.id = ?
	`

	if err := s.DB.Raw(query, id).Scan(&jsonData).Error; err != nil {
		return nil, err
	}

	var response JadwalajarResponse
	if err := json.Unmarshal([]byte(jsonData), &response); err != nil {
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
					'name', g.name,
					'nip', g.nip
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

func (s *JadwalajarService) GetJadwalajarByIDGuru(id uuid.UUID, hari string) ([]JadwalajarAbsenResponse, error) {
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
					'name', g.name,
					'nip', g.nip
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
				'absen_masuk', json_build_object(
					'id', a.id,
					'jam_masuk', a.jam_masuk
				),
				'absen_keluar', json_build_object(
					'id', ak.id,
					'jam_keluar', ak.jam_keluar
				),
				'izin', json_build_object(
					'id', i.id,
					'read', i.read,
					'approval', i.approval
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

			LEFT JOIN LATERAL (
				SELECT a.id, a.jam_masuk
				FROM absen_masuks a
				WHERE a.jadwal_ajar_id = j.id::uuid AND a.tanggal = CURRENT_DATE
				ORDER BY a.jam_masuk LIMIT 1
			) a ON true

			LEFT JOIN LATERAL (
				SELECT ak.id, ak.jam_keluar
				FROM absen_keluars ak
				WHERE ak.absen_masuk_id = a.id::uuid
				ORDER BY ak.jam_keluar LIMIT 1
			) ak ON true

			LEFT JOIN LATERAL (
				SELECT i.id, i.approval, i.read
				FROM izins i
				WHERE i.jadwal_ajar_id = j.id::uuid AND i.tanggal_izin = CURRENT_DATE
				ORDER BY i.updated_at DESC NULLS LAST LIMIT 1
			) i ON true

			%s
			ORDER BY j.jam_mulai
		) sub;
	`, whereClause)

	if err := s.DB.Raw(query, args...).Scan(&jsonData).Error; err != nil {
		return nil, err
	}



	jadwalajarList := []JadwalajarAbsenResponse{}
	if jsonData == nil {
		return jadwalajarList, nil
	}

	if err := json.Unmarshal([]byte(*jsonData), &jadwalajarList); err != nil {
		return nil, err
	}

	return jadwalajarList, nil
}




func (s *JadwalajarService) GetJadwalajarByIDKelas(id uuid.UUID, hari string) ([]JadwalajarAbsenResponse, error) {
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
				'absen_masuk', json_build_object(
					'id', a.id,
					'jam_masuk', a.jam_masuk
				),
				'absen_keluar', json_build_object(
					'id', ak.id,
					'jam_keluar', ak.jam_keluar
				),
				'izin', json_build_object(
					'id', i.id,
					'read', i.read,
					'approval', i.approval
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

			LEFT JOIN LATERAL (
				SELECT a.id, a.jam_masuk
				FROM absen_masuks a
				WHERE a.jadwal_ajar_id = j.id::uuid AND a.tanggal = CURRENT_DATE
				ORDER BY a.jam_masuk
				LIMIT 1
			) a ON true

			LEFT JOIN LATERAL (
				SELECT ak.id, ak.jam_keluar
				FROM absen_keluars ak
				WHERE ak.absen_masuk_id = a.id::uuid
				ORDER BY ak.jam_keluar
				LIMIT 1
			) ak ON true

			LEFT JOIN LATERAL (
				SELECT i.id, i.approval, i.read
				FROM izins i
				WHERE i.jadwal_ajar_id = j.id::uuid AND i.tanggal_izin = CURRENT_DATE
				ORDER BY i.created_at DESC
				LIMIT 1
			) i ON true

			%s
			ORDER BY j.jam_mulai
		) sub;
	`, whereClause)

	if err := s.DB.Raw(query, args...).Scan(&jsonData).Error; err != nil {
		return nil, err
	}

	jadwalajarList := []JadwalajarAbsenResponse{}
	if jsonData == nil {
		return jadwalajarList, nil
	}

	if err := json.Unmarshal([]byte(*jsonData), &jadwalajarList); err != nil {
		return nil, err
	}

	return jadwalajarList, nil
}

