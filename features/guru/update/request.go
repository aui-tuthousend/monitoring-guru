package update

type UpdateGuruRequest struct {
	Name     *string `json:"name,omitempty"`
	Nip      *string `json:"nip,omitempty"`
	Password *string `json:"password,omitempty"`
	Jabatan  *string `json:"jabatan,omitempty"`
}
