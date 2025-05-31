package create

type CreateRuanganResponse struct {
	RuanganID string `json:"ruangan_id"`
	Name      string `json:"nama"`
}

type CreateRuanganResponseWrapper struct {
	Code    int                   `json:"code"`
	Message string                `json:"message"`
	Data    CreateRuanganResponse `json:"data"`
}
