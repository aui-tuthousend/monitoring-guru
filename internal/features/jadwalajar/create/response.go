package create

type CreateJadwalAjarResponse struct {
	ID string `json:"id"`
	GuruID string `json:"guru_id"`
	MapelID string `json:"mapel_id"`
	KelasID string `json:"kelas_id"`
	Hari string `json:"hari"`
	JamMulai string `json:"jam_mulai"`
	JamSelesai string `json:"jam_selesai"`
	LastEditor string `json:"last_editor"`
}

type CreateJadwalAjarResponseWrapper struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Data    CreateJadwalAjarResponse `json:"data"`
}
