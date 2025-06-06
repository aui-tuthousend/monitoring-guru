package kelas

import "monitoring-guru/internal/features/jurusan"

type KelasHandler struct {
	Service        *KelasService
	JurusanService *jurusan.JurusanService
}
