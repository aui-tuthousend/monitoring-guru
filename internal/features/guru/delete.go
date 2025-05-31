package guru

import (

	"github.com/gofiber/fiber/v2"
	e "monitoring-guru/entities"
)

// DeleteGuruHandler godoc
// @Summary Delete a guru
// @Description Delete a guru by ID
// @Tags Guru
// @Security BearerAuth
// @Produce json
// @Param id path string true "Guru ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/guru/{id} [delete]
func (h *GuruHandler) DeleteGuruHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		if err := h.Service.DeleteGuru(id); err != nil {
			return c.Status(500).JSON(e.ErrorResponse[any](500, "Internal server error", nil))
		}

		return c.JSON(fiber.Map{"message": "Guru deleted successfully"})
	}
}