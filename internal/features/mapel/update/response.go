package update

type UpdateMapelResponse struct {
	MapelID   string `json:"mapel_id"`
	JurusanID string `json:"jurusan_id"`
	Name      string `json:"nama"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateMapelResponseWrapper struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Data    UpdateMapelResponse `json:"data"`
}
