package guru

import (
	e "monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
)

// GetAllGuruHandler godoc
// @Summary Get all guru
// @Description Get all guru
// @Tags Guru
// @Security BearerAuth
// @Produce json
// @Success 200 {object} []GuruResponse
// @Failure 500 {object} GuruResponseWrapper
// @Router /api/guru [get]
func (h *GuruHandler) GetAllGuruHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		gurus, err := h.Service.GetAllGuru()
		if err != nil {
			return c.Status(500).JSON(e.ErrorResponse[any](500, "Internal server error", nil))
		}

		var responses []GuruResponse
		for _, guru := range gurus {
			responses = append(responses, GuruResponse{
				ID:      guru.ID.String(),
				NIP:     guru.Nip,
				Nama:    guru.Name,
				Jabatan: guru.Jabatan,
			})
		}

		return c.JSON(e.SuccessResponse(&responses))
	}
}
