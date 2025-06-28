package jadwalajar

import (
	"monitoring-guru/internal/features/absenmasuk"
	"monitoring-guru/internal/features/guru"
	"monitoring-guru/internal/features/kelas"
	"monitoring-guru/internal/features/mapel"
	"monitoring-guru/internal/features/ruangan"
)


type JadwalajarResponse struct {
	ID         string   `json:"id"`
	Guru       *guru.GuruMiniResponse `json:"guru"`
	Mapel      *mapel.MapelMiniResponse   `json:"mapel"`
	Kelas      *kelas.KelasMiniResponse   `json:"kelas"`
	Ruangan    *ruangan.RuanganResponse   `json:"ruangan"`
	Hari       string   `json:"hari"`
	JamMulai   string   `json:"jam_mulai"`
	JamSelesai string   `json:"jam_selesai"`
}

type JadwalajarGuruResponse struct {
	ID         string   `json:"id"`
	Mapel      *mapel.MapelMiniResponse   `json:"mapel"`
	Kelas      *kelas.KelasMiniResponse   `json:"kelas"`
	Ruangan    *ruangan.RuanganResponse   `json:"ruangan"`
	Hari       string   `json:"hari"`
	JamMulai   string   `json:"jam_mulai"`
	JamSelesai string   `json:"jam_selesai"`
	AbsenMasuk *absenmasuk.AbsenMasukMiniResponse `json:"absen_masuk"`
}

type JadwalajarResponseWrapper struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    JadwalajarResponse `json:"data"`
}