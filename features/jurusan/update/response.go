package update

type UpdateJurusanResponse struct {
	JurusanID string `json:"jurusan_id"`
	Name      string `json:"nama"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateJurusanResponseWrapper struct {
	Code    int                   `json:"code"`
	Message string                `json:"message"`
	Data    UpdateJurusanResponse `json:"data"`
}
