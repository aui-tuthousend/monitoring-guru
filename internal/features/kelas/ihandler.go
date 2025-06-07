package kelas

import (
	"monitoring-guru/internal/features/jurusan"
	"monitoring-guru/internal/features/ketuakelas"
)

type KelasHandler struct {
	Service        *KelasService
	JurusanService *jurusan.JurusanService
	KetuaKelasService *ketuakelas.KetuaKelasService
}
