package update

type UpdateMapelRequest struct {
	JurusanID string `json:"jurusan_id" validate:"required,uuid4"`
	Name      string `json:"nama" validate:"required"`
}
