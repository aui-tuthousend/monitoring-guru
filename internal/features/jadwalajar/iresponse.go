package jadwalajar

type JadwalajarResponse struct {
	ID         string `json:"id"`
	Guru       string `json:"guru"`
	Mapel      string `json:"mapel"`
	Kelas      string `json:"kelas"`
	Hari       string    `json:"hari"`
	JamMulai   string         `json:"jam_mulai"`
	JamSelesai string         `json:"jam_selesai"`
	LastEditor string         `json:"last_editor"`
}

type JadwalajarResponseWrapper struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    JadwalajarResponse `json:"data"`
}