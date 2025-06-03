package kelas

import (
	e "monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
)

// GetKelasHandler godoc
// @Summary Get all kelas
// @Description Get all kelas
// @Tags kelas
// @Produce json
// @Success 200 {object} []KelasResponse
// @Failure 500 {object} KelasResponseWrapper
// @Router /api/kelas [get]
func (h *KelasHandler) GetKelasHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		kelasList, err := h.Service.GetAllKelas()
		if err != nil {
			return c.Status(500).JSON(e.ErrorResponse[any](500, "Internal server error", nil))
		}

		return c.JSON(e.SuccessResponse(&kelasList))
	}
}
