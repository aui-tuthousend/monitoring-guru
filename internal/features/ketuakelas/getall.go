package ketuakelas

import (
	e "monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
)

// GetAllKetuaKelasHandler godoc
// @Summary Get all ketua kelas
// @Tags Ketua Kelas
// @Security BearerAuth
// @Produce json
// @Success 200 {object} []KetuaKelasResponse
// @Failure 500 {object} KetuaKelasResponseWrapper
// @Router /api/ketua-kelas [get]
func (h *KetuaKelasHandler) GetAllKetuaKelasHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ketuaKelasList, err := h.Service.GetAllKetuaKelas()
		if err != nil {
			return c.Status(500).JSON(e.ErrorResponse[any](500, "Internal server error", nil))
		}

		var responses []KetuaKelasResponse
		for _, ketua := range ketuaKelasList {
			responses = append(responses, *h.Service.ResponseKetuaKelasMapper(&ketua))
		}

		return c.JSON(e.SuccessResponse(&responses))
	}
}