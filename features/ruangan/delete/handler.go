package delete

import (
	"monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// DeleteRuangan godoc
// @Summary Delete ruangan
// @Description Delete a ruangan by ID
// @Tags ruangan
// @Accept json
// @Produce json
// @Param id path string true "Ruangan ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/ruangan/{id} [delete]
func DeleteRuangan(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		ruanganID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid ID format"})
		}

		var ruangan entities.Ruangan
		if err := db.First(&ruangan, "id = ?", ruanganID).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Ruangan not found"})
		}

		if err := db.Delete(&ruangan).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to delete ruangan"})
		}

		return c.JSON(fiber.Map{
			"code":    200,
			"message": "Ruangan deleted successfully",
		})
	}
}
