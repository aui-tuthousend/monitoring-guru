package statuskelas

import "monitoring-guru/internal/features/kelas"

type StatusKelasResponse struct {
	ID        string `json:"id"`
	Kelas   kelas.KelasStatusResponse `json:"kelas"`
	IsActive  bool   `json:"is_active"`
	Mapel string `json:"mapel"`
	Pengajar string `json:"pengajar"`
	Ruangan string `json:"ruangan"`
}
