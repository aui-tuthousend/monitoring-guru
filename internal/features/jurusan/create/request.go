package create

type CreateJurusanRequest struct {
	Name string `json:"nama" validate:"required"`
}
