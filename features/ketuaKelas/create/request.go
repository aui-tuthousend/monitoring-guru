package create

// CreateKetuaRequestBody
// @Description Create ketua request body
// @Description NISN of the ketua
type CreateKetuaRequest struct {
	// @Description NISN of the ketua
	// @Required true
	// @Example "123456789"
	NISN string `json:"nisn"`
	// @Description Name of the ketua
	// @Required true
	// @Example "John Doe"
	Nama string `json:"nama"`
	// @Description Password of the ketua
	// @Required true
	// @Example "password123"
	// @MinLength 6
	Password string `json:"password"`
}
