package getall

import "monitoring-guru/entities"

type GetAllJadwalAjarResponse struct {
	ID string `json:"id"`
	Guru entities.Guru `json:"guru"`
	Mapel entities.Mapel `json:"mapel"`
	Kelas entities.Kelas `json:"kelas"`
	Hari string `json:"hari"`
	JamMulai string `json:"jam_mulai"`
	JamSelesai string `json:"jam_selesai"`
	LastEditor string `json:"last_editor"`
}

type GetAllJadwalAjarResponseWrapper struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Data    []GetAllJadwalAjarResponse `json:"data"`
}
