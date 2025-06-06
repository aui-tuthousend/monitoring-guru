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
// @Failure 400 {object} MapelResponseWrapper
// @Failure 500 {object} MapelResponseWrapper
// @Router /api/mapel/{id} [get]
func (h *MapelHandler) GetMapelByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		parsedID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, "Invalid UUID format", nil))
		}

		mapel, err := h.Service.GetMapelByID(parsedID)
		if err != nil {
			return c.Status(500).JSON(e.ErrorResponse[any](500, "Mapel not found", nil))
		}

		return c.JSON(e.SuccessResponse(&mapel))
	}
}
