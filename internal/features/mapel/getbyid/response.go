package getbyid

type MapelResponse struct {
	MapelID   string `json:"mapel_id"`
	JurusanID string `json:"jurusan_id"`
	Name      string `json:"nama"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type GetMapelByIDResponseWrapper struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    MapelResponse `json:"data"`
}
