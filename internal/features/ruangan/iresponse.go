package ruangan

type RuanganResponse struct {
	ID string `json:"id" gorm:"column:id"`
	Name      string `json:"name"`
}

type RuanganResponseWrapper struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    RuanganResponse `json:"data"`
}
