package jadwalajar

import (
	"monitoring-guru/internal/features/guru"
	"monitoring-guru/internal/features/kelas"
	"monitoring-guru/internal/features/mapel"
)

type JadwalajarResponse struct {
	ID         string   `json:"id"`
	Guru*       guru.GuruResponse `json:"guru"`
	Mapel*      mapel.MapelResponse   `json:"mapel"`
	Kelas*      kelas.KelasResponse   `json:"kelas"`
	Hari       string   `json:"hari"`
	JamMulai   string   `json:"jam_mulai"`
	JamSelesai string   `json:"jam_selesai"`
	LastEditor string   `json:"last_editor"`
}

type JadwalajarResponseWrapper struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    JadwalajarResponse `json:"data"`
}