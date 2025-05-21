package login

// LoginRequestBody
// @Description Login user request body
type AuthRequest struct {
	// Your Email
	Email string `json:"email"`
	// Your Password
	Password string `json:"password"`
}

// LoginGuruRequestBody
// @Description Login guru request body

type AuthGuruRequest struct {
	// Your NIP
	NIP string `json:"nip"`
	// Your Password
	Password string `json:"password"`
}

// LoginKetuaKelasRequestBody
// @Description Login ketua kelas request body
type AuthKetuaKelasRequest struct {
	// Your NISN
	NISN string `json:"nisn"`
	// Your Password
	Password string `json:"password"`
}
