package mapel

import (
	e "monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetMapelByIDHandler godoc
// @Summary Get mapel by ID
// @Description Get mapel by ID
// @Tags Mapel
// @Security BearerAuth
// @Produce json
// @Param id path string true "Mapel ID"
// @Success 200 {object} MapelResponseWrapper
// @Failure 404 {object} MapelResponseWrapper
// @Router /api/mapel/{id} [get]
func (h *MapelHandler) GetMapelByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		if _, err := uuid.Parse(id); err != nil {
			return c.Status(404).JSON(e.ErrorResponse[any](404, "Invalid UUID format", nil))
		}

		mapel, err := h.Service.GetMapelByID(id)
		if err != nil {
			return c.Status(404).JSON(e.ErrorResponse[any](404, "Mapel not found", nil))
		}

		return c.JSON(e.SuccessResponse(&mapel))
	}
}
