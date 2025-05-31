package update

// UpdateJadwalAjarRequestBody
// @Description Update jadwalajar request body
type UpdateJadwalAjarRequest struct {
	// @Description ID of the jadwalajar
	// @Required true
	// @Example "123456789"
	ID string `json:"id"`
	// @Description Guru ID of the jadwalajar
	// @Required true
	// @Example "123456789"
	GuruID string `json:"guru_id"`
	// @Description Mapel ID of the jadwalajar
	// @Required true
	// @Example "John Doe"
	MapelID string `json:"mapel_id"`
	// @Description Kelas ID of the jadwalajar
	// @Required true
	// @Example "adasd323"
	KelasID string `json:"kelas_id"`
	// @Description Hari of the jadwalajar
	// @Required true
	// @Example "Senin"
	Hari string `json:"hari"`
	// @Description Jam Mulai of the jadwalajar
	// @Required true
	// @Example "08:00"
	JamMulai string `json:"jam_mulai"`
	// @Description Jam Selesai of the jadwalajar
	// @Required true
	// @Example "10:00"
	JamSelesai string `json:"jam_selesai"`
	// @Description Last Editor of the jadwalajar
	// @Example "John Doe"
	LastEditor string `json:"last_editor"`
}
