package ketuakelas

import (
	"github.com/google/uuid"
)

type KetuaKelasResponse struct {
	// @Description ID of the ketua
	// @Required true
	// @Example "123e4567-e89b-12d3-a456-426614174000"
	ID uuid.UUID `json:"id"`
	// @Description NISN of the ketua
	// @Required true
	// @Example "123456789"
	Nisn string `json:"nisn"`
	// @Description Name of the ketua
	// @Required true
	// @Example "John Doe"
	Name string `json:"nama"`
}

type KetuaKelasResponseWrapper struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    KetuaKelasResponse `json:"data"`
}
