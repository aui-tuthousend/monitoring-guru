package getall

type JurusanResponse struct {
	JurusanID string `json:"jurusan_id"`
	Name      string `json:"nama"`
}

type GetAllJurusanResponseWrapper struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    []JurusanResponse `json:"data"`
}
