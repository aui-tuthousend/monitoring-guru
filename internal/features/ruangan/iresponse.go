package ruangan

type RuanganResponse struct {
	RuanganID string `json:"ruangan_id" gorm:"column:id"`
	Name      string `json:"nama"`
}

type RuanganResponseWrapper struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    RuanganResponse `json:"data"`
}
