package getbyidguru

// GetJadwalAjarGuruRequestBody
// @Description Get jadwalajar by ID Guru request body
type GetJadwalAjarGuruRequest struct {
	// @Description Guru ID of the jadwalajar
	// @Required true
	// @Example "123456789"
	GuruID string `json:"guru_id"`
	// @Description Hari of the jadwalajar
	// @Required true
	// @Example "Senin"
	Hari string `json:"hari"`
}
