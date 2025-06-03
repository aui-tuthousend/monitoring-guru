package kelas

import (
	"monitoring-guru/internal/features/ketuakelas"
)

type KelasResponse struct {
	ID   string `json:"id"`
	Nama string `json:"nama"`
	// Jurusan    entities.Jurusan `json:"jurusan"`
	KetuaKelas ketuakelas.KetuaKelasResponse `json:"ketua_kelas"`
	WakilKelas ketuakelas.KetuaKelasResponse `json:"wakil_kelas"`
	IsActive   bool                          `json:"is_active"`
}

type KelasResponseWrapper struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    KelasResponse `json:"data"`
}
