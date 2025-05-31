package update

type UpdateRuanganResponse struct {
	RuanganID string `json:"ruangan_id"`
	Name      string `json:"nama"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateRuanganResponseWrapper struct {
	Code    int                   `json:"code"`
	Message string                `json:"message"`
	Data    UpdateRuanganResponse `json:"data"`
}
