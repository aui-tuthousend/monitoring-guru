package getall

type RuanganResponse struct {
	RuanganID string `json:"ruangan_id"`
	Name      string `json:"nama"`
}

type GetAllRuanganResponseWrapper struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    []RuanganResponse `json:"data"`
}
