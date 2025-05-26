package getbyid

type RuanganResponse struct {
	RuanganID string `json:"ruangan_id"`
	Name      string `json:"nama"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type GetRuanganByIDResponseWrapper struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    RuanganResponse `json:"data"`
}
