package websocket

import (
	"encoding/json"
	"log"
	"monitoring-guru/entities"
	"monitoring-guru/internal/features/absenkeluar"
	"monitoring-guru/internal/features/absenmasuk"
	"monitoring-guru/internal/features/izin"
	"monitoring-guru/internal/features/jadwalajar"
	"monitoring-guru/utils"
	"time"

	"github.com/google/uuid"
)

type WebsocketService struct {
	JadwalajarService *jadwalajar.JadwalajarService
	AbsenMasukService *absenmasuk.AbsenMasukService
	AbsenKeluarService *absenkeluar.AbsenKeluarService
	IzinService *izin.IzinService
}

func (s *WebsocketService) CreateAbsenMasuk(data json.RawMessage) bool {
	var payload struct {
		Id            string `json:"id"`
		IsActive      bool   `json:"is_active"`
		JadwalajarID  string `json:"jadwalajar_id"`
		Mapel         string `json:"mapel"`
		Pengajar      string `json:"pengajar"`
		Ruangan       string `json:"ruangan"`
	}

	if err := json.Unmarshal(data, &payload); err != nil {
		log.Println("Error unmarshalling payload:", err)
		return false
	}
	if payload.Id == "" {
		log.Println("Payload missing ID")
		return false
	}

	log.Printf("Parsed payload: %+v\n", payload)

	jadwalajar, err := s.JadwalajarService.GetJadwalajarByID(payload.JadwalajarID)
	if err != nil {
		log.Printf("Failed to get JadwalAjar: %v", err)
		BroadcastToGroup("admin", "Failed")
		return false
	}

	errr := s.JadwalajarService.DB.Model(&entities.StatusKelas{}).
		Where("kelas_id = ?", payload.Id).
		Updates(map[string]interface{}{
			"is_active": payload.IsActive,
			"mapel":     jadwalajar.Mapel.Name,
			"pengajar":  jadwalajar.Guru.Name,
			"ruangan":   jadwalajar.Ruangan.Name,
		}).Error

	if errr != nil {
		log.Printf("Failed to update DB: %v", err)
		BroadcastToGroup("admin", "Failed")
		return false
	}


	jadwalajarID, err := utils.ParseUUID(jadwalajar.ID)
	if err != nil {
		log.Println("Invalid jadwalajar ID:", err)
		return false
	}
	guruID, err := utils.ParseUUID(jadwalajar.Guru.ID)
	if err != nil {
		log.Println("Invalid guru ID:", err)
		return false
	}
	kelasID, err := utils.ParseUUID(jadwalajar.Kelas.ID)
	if err != nil {
		log.Println("Invalid kelas ID:", err)
		return false
	}
	ruanganID, err := utils.ParseUUID(jadwalajar.Ruangan.ID)
	if err != nil {
		log.Println("Invalid ruangan ID:", err)
		return false
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	absenMasuk := &entities.AbsenMasuk{
		ID:           uuid.New(),
		GuruID:       guruID,
		JadwalAjarID: jadwalajarID,
		KelasID:      kelasID,
		RuanganID:    ruanganID,
		Tanggal:      time.Now().In(loc),
		JamMasuk:     time.Now().In(loc).Format("15:04"),
	}

	_, err = s.AbsenMasukService.CreateAbsenMasuk(absenMasuk)
	if err != nil {
		log.Printf("Failed to create AbsenMasuk: %v", err)
		BroadcastToGroup("admin", "Failed")
		return false
	}

	payload.Pengajar = jadwalajar.Guru.Name
	payload.Ruangan = jadwalajar.Ruangan.Name
	payload.Mapel = jadwalajar.Mapel.Name

	response, _ := json.Marshal(struct {
		Type    string      `json:"type"`
		Payload interface{} `json:"payload"`
	}{
		Type:    "update-kelas",
		Payload: payload,
	})

	siswaMessage, _ := json.Marshal(struct {
		Type    string `json:"type"`
		Payload string `json:"payload"`
	}{
		Type:    "absen-success",
		Payload: "Absen berhasil",
	})

	BroadcastToGroup("admin", string(response))
	BroadcastToGroup("siswa", string(siswaMessage))
	return true
}





func (s *WebsocketService) CreateAbsenKeluar(data json.RawMessage) bool {
	var payload struct {
		Id            string `json:"id"`
		IsActive      bool   `json:"is_active"`
		// JadwalajarID  string `json:"jadwalajar_id"`
		AbsenMasukID  string `json:"absen_masuk_id"`
		Mapel         string `json:"mapel"`
		Pengajar      string `json:"pengajar"`
		Ruangan       string `json:"ruangan"`
	}

	if err := json.Unmarshal(data, &payload); err != nil {
		log.Println("Error unmarshalling payload:", err)
		return false
	}
	if payload.Id == "" {
		log.Println("Payload missing ID")
		return false
	}

	log.Printf("Parsed payload: %+v\n", payload)

	errr := s.JadwalajarService.DB.Model(&entities.StatusKelas{}).
		Where("kelas_id = ?", payload.Id).
		Updates(map[string]interface{}{
			"is_active": payload.IsActive,
			"mapel":     "-",
			"pengajar":  "-",
			"ruangan":   "-",
		}).Error

	if errr != nil {
		log.Printf("Failed to update DB: %v", errr)
		BroadcastToGroup("admin", "Failed")
		return false
	}

	absenMasukID, err := utils.ParseUUID(payload.AbsenMasukID)
	if err != nil {
		log.Println("Invalid absen masuk ID:", err)
		return false
	}


	loc, _ := time.LoadLocation("Asia/Jakarta")
	absenKeluar := &entities.AbsenKeluar{
		ID:           uuid.New(),
		AbsenMasukID: absenMasukID,
		JamKeluar:    time.Now().In(loc).Format("15:04"),
		Status:       "Keluar",
	}

	err = s.AbsenKeluarService.CreateAbsenKeluar(absenKeluar)
	if err != nil {
		log.Printf("Failed to create AbsenKeluar: %v", err)
		BroadcastToGroup("admin", "Failed")
		return false
	}

	payload.Pengajar = "-"
	payload.Ruangan = "-"
	payload.Mapel = "-"
	payload.IsActive = false

	response, _ := json.Marshal(struct {
		Type    string      `json:"type"`
		Payload interface{} `json:"payload"`
	}{
		Type:    "update-kelas",
		Payload: payload,
	})
	siswaMessage, _ := json.Marshal(struct {
		Type    string `json:"type"`
		Payload string `json:"payload"`
	}{
		Type:    "absen-success",
		Payload: "Absen berhasil",
	})

	BroadcastToGroup("admin", string(response))
	BroadcastToGroup("siswa", string(siswaMessage))
	return true
}



func (s *WebsocketService) CreateIzin(data json.RawMessage) bool {
	var payload struct {
		Id string `json:"id"`
		Judul string `json:"judul"`
		Pesan string `json:"pesan"`
		JadwalajarID  string `json:"jadwalajar_id"`
		JamIzin string `json:"jam_izin"`
		TanggalIzin string `json:"tanggal_izin"`
		Read bool `json:"read"`
		Approval bool `json:"approval"`
		Guru string `json:"guru"`
	}

	if err := json.Unmarshal(data, &payload); err != nil {
		log.Println("Error unmarshalling payload:", err)
		return false
	}

	log.Printf("Parsed payload: %+v\n", payload)

	loc, _ := time.LoadLocation("Asia/Jakarta")
	izinEntity := entities.Izin{
		ID:           uuid.New(),
		// GuruID:       uuid.MustParse(req.GuruID),
		Judul:       payload.Judul,
		Pesan:        payload.Pesan,
		JadwalAjarID: uuid.MustParse(payload.JadwalajarID),
		JamIzin: time.Now().In(loc).Format("15:04"),
		TanggalIzin:  time.Now().In(loc),
		Approval:     false,
		Read: false,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err := s.IzinService.CreateIzin(&izinEntity)
	if err != nil {
		log.Printf("Failed to create Izin: %v", err)
		BroadcastToGroup("admin", "Failed")
		return false
	}

	izin, err := s.IzinService.GetIzinByID(izinEntity.ID.String())
	if err != nil {
		log.Printf("Failed to get Izin: %v", err)
		BroadcastToGroup("admin", "Failed")
		return false
	}

	response, _ := json.Marshal(struct {
		Type    string      `json:"type"`
		Payload interface{} `json:"payload"`
	}{
		Type:    "izin-masuk",
		Payload: izin,
	})

	BroadcastToGroup("admin", string(response))
	return true
}

func (s *WebsocketService) HandleIzin(data json.RawMessage) bool {
	var payload struct {
		Id string `json:"id"`
		Status bool `json:"status"`
	}

	if err := json.Unmarshal(data, &payload); err != nil {
		log.Println("Error unmarshalling payload:", err)
		return false
	}

	log.Printf("Parsed payload: %+v\n", payload)

	izin, err := s.IzinService.GetIzin(payload.Id)
	if err != nil {
		log.Printf("Failed to get Izin: %v", err)
		BroadcastToGroup("admin", "Failed")
		return false
	}

	izin.Approval = payload.Status
	izin.Read = true
	err = s.IzinService.UpdateIzin(izin)
	if err != nil {
		log.Printf("Failed to update Izin: %v", err)
		BroadcastToGroup("admin", "Failed")
		return false
	}

	var statusIzin string
	if payload.Status {
		statusIzin = "disetujui"
	} else {
		statusIzin = "ditolak"
	}

	jadwalajar, _ := s.JadwalajarService.GetJadwalajarByID(izin.JadwalAjarID.String())
	guruID := "user-" + jadwalajar.Guru.Nip

	response, _ := json.Marshal(struct {
		Type    string      `json:"type"`
		Payload interface{} `json:"payload"`
	}{
		Type:    "handle-izin",
		Payload: "Izin telah " + statusIzin,
	})
	SendToUserInGroup("guru", guruID, string(response))
	return true
}