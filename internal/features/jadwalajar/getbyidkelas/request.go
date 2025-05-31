package getbyidkelas

// GetJadwalAjarKelasRequestBody
// @Description Get jadwalajar by ID Kelas request body
type GetJadwalAjarKelasRequest struct {
	// @Description Kelas ID of the jadwalajar
	// @Required true
	// @Example "123456789"
	KelasID string `json:"kelas_id"`
	// @Description Hari of the jadwalajar
	// @Required true
	// @Example "Senin"
	Hari string `json:"hari"`
}
