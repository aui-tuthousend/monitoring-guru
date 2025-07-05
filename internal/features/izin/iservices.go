package izin

import (
	"encoding/json"
	e "monitoring-guru/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IzinService struct {
	DB *gorm.DB
}

func (s *IzinService) CreateIzin(izin *e.Izin) error {
	return s.DB.Create(izin).Error
}

func (s *IzinService) GetIzin(id string) (*e.Izin, error) {
	var izin e.Izin
	if err := s.DB.Where("id = ?", id).First(&izin).Error; err != nil {
		return nil, err
	}
	return &izin, nil
}

func (s *IzinService) ResponseIzinMapper(izin *e.Izin) *IzinResponse {
	return &IzinResponse{
		ID:           izin.ID.String(),
		Judul:       izin.Judul,
		// GuruID:       izin.GuruID.String(),
		// JadwalAjarID: izin.JadwalAjarID.String(),
		TanggalIzin:  izin.TanggalIzin.Format("2006-01-02"), // Format YYYY-MM-DD
		JamIzin: izin.JamIzin,
		Pesan:        izin.Pesan,
		Read: izin.Read,
		Approval:     izin.Approval,
	}
}

func (s *IzinService) GetAllIzin() ([]IzinResponse, error) {
	var jsonData *string
	query := `
		SELECT json_agg(result)
		FROM (
			SELECT json_build_object(
				'id', i.id,
				'judul', i.judul,
				'pesan', i.pesan,
				'guru', g.name,
				'mapel', m.name,
				'jam_mulai', j.jam_mulai,
				'jam_selesai', j.jam_selesai,
				'tanggal_izin', i.tanggal_izin,
				'jam_izin', i.jam_izin,
				'read', i.read,
				'approval', i.approval
			) AS result
			FROM izins i
			JOIN jadwal_ajars j ON j.id = i.jadwal_ajar_id::uuid
			JOIN gurus g ON g.id = j.guru_id::uuid
			JOIN mapels m ON m.id = j.mapel_id::uuid
			WHERE i.read = false
			ORDER BY i.jam_izin DESC
		) sub;
	`
	if err := s.DB.Raw(query).Scan(&jsonData).Error; err != nil {
		return nil, err
	}
	izinList := []IzinResponse{}
	if jsonData == nil {
		return izinList, nil
	}
	if err := json.Unmarshal([]byte(*jsonData), &izinList); err != nil {
		return nil, err
	}
	return izinList, nil
}

func (s *IzinService) GetAllIzinGuru(nip string) ([]IzinResponse, error) {
	var jsonData *string
	query := `
		SELECT json_agg(result)
		FROM (
			SELECT json_build_object(
				'id', i.id,
				'judul', i.judul,
				'pesan', i.pesan,
				'guru', g.name,
				'mapel', m.name,
				'jam_mulai', j.jam_mulai,
				'jam_selesai', j.jam_selesai,
				'tanggal_izin', i.tanggal_izin,
				'jam_izin', i.jam_izin,
				'read', i.read,
				'approval', i.approval
			) AS result
			FROM izins i
			JOIN jadwal_ajars j ON j.id = i.jadwal_ajar_id::uuid
			JOIN gurus g ON g.id = j.guru_id::uuid
			JOIN mapels m ON m.id = j.mapel_id::uuid
			WHERE g.nip = ? AND i.tanggal_izin = CURRENT_DATE
			ORDER BY i.jam_izin DESC
		) sub;
	`
	if err := s.DB.Raw(query, nip).Scan(&jsonData).Error; err != nil {
		return nil, err
	}
	izinList := []IzinResponse{}
	if jsonData == nil {
		return izinList, nil
	}
	if err := json.Unmarshal([]byte(*jsonData), &izinList); err != nil {
		return nil, err
	}
	return izinList, nil
}

func (s *IzinService) GetAllIzinKelas(kelasID uuid.UUID) ([]IzinResponse, error) {
	var jsonData *string
	query := `
		SELECT json_agg(result)
		FROM (
			SELECT json_build_object(
				'id', i.id,
				'judul', i.judul,
				'pesan', i.pesan,
				'guru', g.name,
				'mapel', m.name,
				'jam_mulai', j.jam_mulai,
				'jam_selesai', j.jam_selesai,
				'tanggal_izin', i.tanggal_izin,
				'jam_izin', i.jam_izin,
				'read', i.read,
				'approval', i.approval
			) AS result
			FROM izins i
			JOIN jadwal_ajars j ON j.id = i.jadwal_ajar_id::uuid
			JOIN gurus g ON g.id = j.guru_id::uuid
			JOIN mapels m ON m.id = j.mapel_id::uuid
			WHERE j.kelas_id = ?::uuid AND i.tanggal_izin = CURRENT_DATE
			AND i.approval = true
			ORDER BY i.jam_izin DESC
		) sub;
	`
	if err := s.DB.Raw(query, kelasID).Scan(&jsonData).Error; err != nil {
		return nil, err
	}
	izinList := []IzinResponse{}
	if jsonData == nil {
		return izinList, nil
	}
	if err := json.Unmarshal([]byte(*jsonData), &izinList); err != nil {
		return nil, err
	}
	return izinList, nil
}


func (s *IzinService) GetIzinByID(id string) (*IzinResponse, error) {
	var jsonData string

	query := `
		SELECT json_build_object(
			'id', i.id,
			'judul', i.judul,
			'pesan', i.pesan,
			'guru', g.name,
			'mapel', m.name,
			'jam_mulai', j.jam_mulai,
			'jam_selesai', j.jam_selesai,
			'tanggal_izin', i.tanggal_izin,
			'jam_izin', i.jam_izin,
			'read', i.read,
			'approval', i.approval
		)
		FROM izins i
		JOIN jadwal_ajars j ON j.id = i.jadwal_ajar_id::uuid
		JOIN gurus g ON g.id = j.guru_id::uuid
		JOIN mapels m ON m.id = j.mapel_id::uuid
		WHERE i.id = ?::uuid
	`

	if err := s.DB.Raw(query, id).Scan(&jsonData).Error; err != nil {
		return nil, err
	}

	if jsonData == "" {
		return nil, nil // or return custom not found error
	}

	var izin IzinResponse
	if err := json.Unmarshal([]byte(jsonData), &izin); err != nil {
		return nil, err
	}

	return &izin, nil
}

func (s *IzinService) IsIzinToday(jadwal_id uuid.UUID) (bool, error) {
	var izin e.Izin
	if err := s.DB.Where("jadwal_ajar_id = ?::uuid AND tanggal_izin = CURRENT_DATE", jadwal_id).First(&izin).Error; err != nil {
		return false, err
	}
	return true, nil
}
	


func (s *IzinService) UpdateIzin(izin *e.Izin) error {
	return s.DB.Save(izin).Error
}

func (s *IzinService) DeleteIzin(id string) error {
	var izin e.Izin
	if err := s.DB.Where("id = ?", id).First(&izin).Error; err != nil {
		return err
	}
	return s.DB.Delete(&izin).Error
}
