package delete

import (
	"monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// DeleteJurusan godoc
// @Summary Delete jurusan
// @Description Delete jurusan by ID
// @Tags jurusan
// @Accept json
// @Produce json
// @Param id path string true "Jurusan ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/jurusan/{id} [delete]
func DeleteJurusan(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		jurusanID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid ID format",
			})
		}

		var jurusan entities.Jurusan
		if err := db.First(&jurusan, "id = ?", jurusanID).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Jurusan not found",
			})
		}

		if err := db.Delete(&jurusan).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete jurusan",
			})
		}

		return c.JSON(fiber.Map{
			"code":    fiber.StatusOK,
			"message": "Jurusan deleted successfully",
		})
	}
}
