package getbyidguru

import "monitoring-guru/entities"

type GetByIDGuruResponse struct {
	ID string `json:"id"`
	Mapel entities.Mapel `json:"mapel"`
	Kelas entities.Kelas `json:"kelas"`
	Hari string `json:"hari"`
	JamMulai string `json:"jam_mulai"`
	JamSelesai string `json:"jam_selesai"`
	LastEditor string `json:"last_editor"`
}

type GetByIDGuruResponseWrapper struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Data    []GetByIDGuruResponse `json:"data"`
}
