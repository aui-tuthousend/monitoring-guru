package getall

import "github.com/google/uuid"

type GetAllKetuaKelasResponse struct {
	// @Description ID of the ketua
	// @Required true
	// @Example "123e4567-e89b-12d3-a456-426614174000"
	ID uuid.UUID `json:"id"`
	// @Description NISN of the ketua
	// @Required true
	// @Example "123456789"
	NISN string `json:"nisn"`
	// @Description Name of the ketua
	// @Required true
	// @Example "John Doe"
	Name string `json:"nama"`
}
