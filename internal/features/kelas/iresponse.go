package kelas

import (
	"monitoring-guru/internal/features/jurusan"
	"monitoring-guru/internal/features/ketuakelas"
)

type KelasResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Jurusan *jurusan.JurusanResponse `json:"jurusan"`
	KetuaKelas *ketuakelas.KetuaKelasResponse `json:"ketua_kelas"`
	IsActive   bool                          `json:"is_active"`
}

type KelasMiniResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type KelasResponseWrapper struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    KelasResponse `json:"data"`
}
