package jurusan

type JurusanResponse struct {
	JurusanID string `json:"jurusan_id"`
	Name      string `json:"nama"`
}

type JurusanResponseWrapper struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    JurusanResponse `json:"data"`
}
