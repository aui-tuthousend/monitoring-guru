package jadwalajar

import (
	"monitoring-guru/internal/features/guru"
	"monitoring-guru/internal/features/kelas"
	"monitoring-guru/internal/features/mapel"
)

type JadwalajarHandler struct {
	Service     *JadwalajarService
	GuruService *guru.GuruService
	MapelService *mapel.MapelService
	KelasService *kelas.KelasService
}