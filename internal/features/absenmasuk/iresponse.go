package absenmasuk

import (
	"monitoring-guru/internal/features/guru"
	"monitoring-guru/internal/features/jadwalajar"
	"monitoring-guru/internal/features/ruangan"
)

type CreateAbsenMasukResponse struct {
	KelasID  string `json:"kelas_id"`
	Ruangan  ruangan.RuanganResponse `json:"ruangan"`
	IsActive bool   `json:"is_active"`
}

type GetAbsenMasukResponse struct {
	ID        string `json:"id"`
	Guru      guru.GuruResponse `json:"guru"`
	JadwalAjar jadwalajar.JadwalajarResponse `json:"jadwal_ajar"`
	Ruangan   ruangan.RuanganResponse `json:"ruangan"`
	Tanggal   string `json:"tanggal"`
	JamMasuk  string `json:"jam_masuk"`
}

type GetAbsenMasukResponseWrapper struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []GetAbsenMasukResponse `json:"data"`
}