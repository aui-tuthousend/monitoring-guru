package create

type CreateJurusanResponse struct {
	JurusanID string `json:"jurusan_id"`
	Name      string `json:"nama"`
}

type CreateJurusanResponseWrapper struct {
	Code    int                   `json:"code"`
	Message string                `json:"message"`
	Data    CreateJurusanResponse `json:"data"`
}
