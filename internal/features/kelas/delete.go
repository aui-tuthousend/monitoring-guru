package kelas

import (
	e "monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
)

// DeleteKelasHandler godoc
// @Summary Delete a kelas
// @Description Delete a kelas by ID
// @Tags Kelas
// @Security BearerAuth
// @Produce json
// @Param id path string true "Kelas ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/kelas/{id} [delete]
func (h *KelasHandler) DeleteKelasHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		if err := h.Service.DeleteKelas(id); err != nil {
			return c.Status(500).JSON(e.ErrorResponse[any](500, "Internal server error", nil))
		}

		return c.JSON(fiber.Map{"message": "Kelas deleted successfully"})
	}
}
