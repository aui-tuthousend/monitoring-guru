package create

// CreateGuruRequestBody
// @Description Create guru request body
type CreateGuruRequest struct {
	// @Description NIP of the guru
	// @Required true
	// @Example "123456789"
	NIP string `json:"nip"`
	// @Description Name of the guru
	// @Required true
	// @Example "John Doe"
	Nama string `json:"nama"`
	// @Description Password of the guru
	// @Required true
	// @Example "password123"
	// @MinLength 6
	Password string `json:"password"`
	// @Description Jabatan of the guru
	// @Required true
	// @Enum "guru" "kepala_sekolah"
	// @Example "guru"
	Jabatan string `json:"jabatan"`
}
