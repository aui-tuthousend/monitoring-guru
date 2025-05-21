package delete

import (
	"monitoring-guru/infrastructure/repositories/guru"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// DeleteGuruHandler godoc
// @Summary Delete a guru
// @Description Delete a guru by ID
// @Tags guru
// @Security BearerAuth
// @Produce json
// @Param id path string true "Guru ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/guru/{id} [delete]
func DeleteGuruHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		if err := guru.DeleteGuru(db, id); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
		}

		return c.JSON(fiber.Map{"message": "Guru deleted successfully"})
	}
}
