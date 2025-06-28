package izin

type IzinResponse struct {
	ID           string `json:"id"`
	GuruID       string `json:"guru_id"`
	JadwalAjarID string `json:"jadwal_ajar_id"`
	TanggalIzin  string `json:"tanggal_izin"`
	Pesan        string `json:"pesan"`
	Approval     bool   `json:"approval"`
}

type IzinResponseWrapper struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    IzinResponse `json:"data"`
}
