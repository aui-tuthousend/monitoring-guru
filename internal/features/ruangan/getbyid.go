package ruangan

import (
	e "monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetRuanganByIDHandler godoc
// @Summary Get ruangan by ID
// @Description Ambil data ruangan berdasarkan ID
// @Tags Ruangan
// @Security     BearerAuth
// @Produce json
// @Param id path string true "Ruangan ID"
// @Success 200 {object} RuanganResponseWrapper
// @Failure 404 {object} RuanganResponseWrapper
// @Router /api/ruangan/{id} [get]
func (h *RuanganHandler) GetRuanganByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		if _, err := uuid.Parse(id); err != nil {
			return c.Status(fiber.StatusNotFound).JSON(e.ErrorResponse[any](404, "Invalid UUID format", nil))
		}

		ruangan, err := h.Service.GetRuanganByID(id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(e.ErrorResponse[any](404, "Ruangan not found", nil))
		}

		return c.JSON(e.SuccessResponse(&ruangan))
	}
}
