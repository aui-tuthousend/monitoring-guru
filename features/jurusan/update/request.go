package update

type UpdateJurusanRequest struct {
	Name string `json:"nama" validate:"required"`
}
