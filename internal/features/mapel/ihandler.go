package mapel

import "monitoring-guru/internal/features/jurusan"

type MapelHandler struct {
	Service        *MapelService
	JurusanService *jurusan.JurusanService
}
