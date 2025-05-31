package getbyid

import "monitoring-guru/entities"

type JadwalajarResponse struct {
	ID string `json:"id"`
	Guru entities.Guru `json:"guru"`
	Mapel entities.Mapel `json:"mapel"`
	Kelas entities.Kelas `json:"kelas"`
	Hari string `json:"hari"`
	JamMulai string `json:"jam_mulai"`
	JamSelesai string `json:"jam_selesai"`
}

type GetJadwalajarByIDResponseWrapper struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    JadwalajarResponse `json:"data"`
}
