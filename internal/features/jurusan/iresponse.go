package jurusan

type JurusanResponse struct {
	ID string `json:"id"`
	Name      string `json:"name"`
	KodeJurusan string `json:"kode_jurusan"`
}

type JurusanResponseWrapper struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    JurusanResponse `json:"data"`
}
