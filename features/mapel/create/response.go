package create

type CreateMapelResponse struct {
	MapelID   string `json:"mapel_id"`
	JurusanID string `json:"jurusan_id"`
	Name      string `json:"nama"`
}

type CreateMapelResponseWrapper struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Data    CreateMapelResponse `json:"data"`
}
