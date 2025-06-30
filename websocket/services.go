package websocket

import (
	"encoding/json"
	"log"
	"monitoring-guru/entities"
	"monitoring-guru/internal/features/absenkeluar"
	"monitoring-guru/internal/features/absenmasuk"
	"monitoring-guru/internal/features/jadwalajar"
	"monitoring-guru/utils"
	"time"

	"github.com/google/uuid"
)

type WebsocketService struct {
	JadwalajarService *jadwalajar.JadwalajarService
	AbsenMasukService *absenmasuk.AbsenMasukService
	AbsenKeluarService *absenkeluar.AbsenKeluarService
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
		BroadcastToAll("Failed")
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
		BroadcastToAll("Failed")
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
		BroadcastToAll("Failed")
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

	BroadcastToAll(string(response))
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
		BroadcastToAll("Failed")
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
		BroadcastToAll("Failed")
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

	BroadcastToAll(string(response))
	return true
}
