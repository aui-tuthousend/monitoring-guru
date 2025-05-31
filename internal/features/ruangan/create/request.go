package create

type CreateRuanganRequest struct {
	Name string `json:"nama" validate:"required"`
}
