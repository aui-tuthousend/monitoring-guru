package ruangan

import (
	e "monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
)

// GetAllRuanganHandler godoc
// @Summary Get all ruangan
// @Description Get semua data ruangan
// @Tags Ruangan
// @Security     BearerAuth
// @Produce json
// @Success 200 {object} []RuanganResponse
// @Failure 500 {object} RuanganResponseWrapper
// @Router /api/ruangan [get]
func (h *RuanganHandler) GetAllRuangan() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ruanganList, err := h.Service.GetAllRuangan()
		if err != nil {
			return c.Status(500).JSON(e.ErrorResponse[any](500, "Internal server error", nil))
		}

		return c.JSON(e.SuccessResponse(&ruanganList))
	}
}
