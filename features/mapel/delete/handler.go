package delete

import (
	"monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// DeleteMapel godoc
// @Summary Delete mapel
// @Description Delete a mapel by ID
// @Tags mapel
// @Accept json
// @Produce json
// @Param id path string true "Mapel ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/mapel/{id} [delete]
func DeleteMapel(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		mapelID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
		}

		var mapel entities.Mapel
		result := db.First(&mapel, "id = ?", mapelID)
		if result.Error != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Mapel not found"})
		}

		if err := db.Delete(&mapel).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete mapel"})
		}

		return c.JSON(fiber.Map{
			"code":    fiber.StatusOK,
			"message": "Mapel deleted successfully",
		})
	}
}
