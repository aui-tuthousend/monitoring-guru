package update

type UpdateRuanganRequest struct {
	Name string `json:"nama" validate:"required"`
}
