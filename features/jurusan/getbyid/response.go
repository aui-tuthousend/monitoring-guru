package getbyid

type JurusanResponse struct {
	JurusanID string `json:"jurusan_id"`
	Name      string `json:"nama"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type GetJurusanResponseWrapper struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    JurusanResponse `json:"data"`
}
