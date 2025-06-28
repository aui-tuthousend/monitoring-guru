package izin

import (
	e "monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
)

// DeleteIzinHandler godoc
// @Summary Delete izin
// @Description Delete izin by ID
// @Tags Izin
// @Security BearerAuth
// @Produce json
// @Param id path string true "Izin ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/izin/{id} [delete]
func (h *IzinHandler) DeleteIzinHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		if err := h.Service.DeleteIzin(id); err != nil {
			return c.Status(500).JSON(e.ErrorResponse[any](500, "Internal server error", nil))
		}

		return c.JSON(fiber.Map{"message": "Izin deleted successfully"})
	}
}
