package getall

import "github.com/google/uuid"

// GetAllGuruResponse godoc
// @Description Get all guru response
type GetAllGuruResponse struct {
	// ID of the guru
	// @example 1
	ID uuid.UUID `json:"id"`
	// NIP of the guru
	// @example 1234567890
	NIP string `json:"nip"`
	// Name of the guru
	// @example John Doe
	Name string `json:"name"`
}
