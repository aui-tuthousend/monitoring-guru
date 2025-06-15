package ketuakelas

import (
	e "monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
)

// GetUnsignedKetuaKelasHandler godoc
// @Summary Get all unsigned ketua kelas
// @Tags Ketua Kelas
// @Security BearerAuth
// @Produce json
// @Success 200 {object} []KetuaKelasResponse
// @Failure 500 {object} KetuaKelasResponseWrapper
// @Router /api/ketua-kelas/unsigned [get]
func (h *KetuaKelasHandler) GetUnsignedKetuaKelasHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ketuaKelasList, err := h.Service.GetUnsignedKetuaKelas()
		if err != nil {
			return c.Status(500).JSON(e.ErrorResponse[any](500, "Internal server error", nil))
		}

		return c.JSON(e.SuccessResponse(&ketuaKelasList))
	}
}