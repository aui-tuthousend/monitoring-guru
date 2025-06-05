package mapel

import (
	e "monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// DeleteMapelHandler godoc
// @Summary Delete mapel
// @Description Delete mapel by ID
// @Tags mapel
// @Accept json
// @Produce json
// @Param id path string true "Mapel ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/mapel/{id} [delete]
func (h *MapelHandler) DeleteMapel() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		mapelID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(e.ErrorResponse[any](400, "Invalid ID format", nil))
		}

		if err := h.Service.DeleteMapel(mapelID.String()); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(e.ErrorResponse[any](500, "Failed to delete mapel", nil))
		}

		return c.JSON(fiber.Map{
			"code":    fiber.StatusOK,
			"message": "Mapel deleted successfully",
		})
	}
}
