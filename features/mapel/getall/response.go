package getall

type MapelResponse struct {
	MapelID   string `json:"mapel_id"`
	JurusanID string `json:"jurusan_id"`
	Name      string `json:"nama"`
}

type GetAllMapelResponseWrapper struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    []MapelResponse `json:"data"`
}
