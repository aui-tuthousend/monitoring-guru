package ruangan

import (
	e "monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
)

// GetAllRuanganHandler godoc
// @Summary Get all ruangan
// @Description Get semua data ruangan
// @Tags ruangan
// @Produce json
// @Success 200 {object} []RuanganResponse
// @Failure 500 {object} ErrorResponseWrapper
// @Router /api/ruangan [get]
func (h *RuanganHandler) GetAllRuangan() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ruanganList, err := h.Service.GetAllRuangan()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(e.ErrorResponse[any](500, "Gagal mengambil data ruangan", nil))
		}

		return c.JSON(e.SuccessResponse(&ruanganList))
	}
}
