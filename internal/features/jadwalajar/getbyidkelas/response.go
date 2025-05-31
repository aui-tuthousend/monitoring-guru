package getbyidkelas

import "monitoring-guru/entities"

type GetByIDKelasResponse struct {
	ID string `json:"id"`
	Guru entities.Guru `json:"guru"`
	Mapel entities.Mapel `json:"mapel"`
	Hari string `json:"hari"`
	JamMulai string `json:"jam_mulai"`
	JamSelesai string `json:"jam_selesai"`
	LastEditor string `json:"last_editor"`
}

type GetByIDKelasResponseWrapper struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Data    []GetByIDKelasResponse `json:"data"`
}
