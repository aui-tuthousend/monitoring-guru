package izin

type IzinResponse struct {
	ID           string `json:"id"`
	Judul       string `json:"judul"`
	Pesan        string `json:"pesan"`
	// JadwalAjarID string `json:"jadwal_ajar_id"`
	Guru string `json:"guru"`
	Mapel string `json:"mapel"`
	JamMulai   string   `json:"jam_mulai"`
	JamSelesai string   `json:"jam_selesai"`
	TanggalIzin  string `json:"tanggal_izin"`
	JamIzin      string `json:"jam_izin"`
	Read		bool	`json:"read"`
	Approval     bool   `json:"approval"`
}

type IzinResponseWrapper struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    IzinResponse `json:"data"`
}
