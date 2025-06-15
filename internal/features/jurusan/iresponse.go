package jurusan

type JurusanResponse struct {
	ID string `json:"id"`
	Name      string `json:"name"`
}

type JurusanResponseWrapper struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    JurusanResponse `json:"data"`
}
