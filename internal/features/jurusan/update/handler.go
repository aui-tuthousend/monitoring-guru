package update

import (
	"monitoring-guru/entities"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UpdateJurusan godoc
// @Summary Update jurusan
// @Description Update jurusan data
// @Tags jurusan
// @Accept json
// @Produce json
// @Param id path string true "Jurusan ID"
// @Param request body UpdateJurusanRequest true "Jurusan data"
// @Success 200 {object} UpdateJurusanResponseWrapper
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/jurusan/{id} [put]
func UpdateJurusan(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		jurusanID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid ID format",
			})
		}

		var req UpdateJurusanRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		if req.Name == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Nama jurusan is required",
			})
		}

		var jurusan entities.Jurusan
		if err := db.First(&jurusan, "id = ?", jurusanID).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Jurusan not found",
			})
		}

		jurusan.Name = req.Name
		jurusan.UpdatedAt = time.Now()

		if err := db.Save(&jurusan).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update jurusan",
			})
		}

		return c.JSON(UpdateJurusanResponseWrapper{
			Code:    fiber.StatusOK,
			Message: "Jurusan updated successfully",
			Data: UpdateJurusanResponse{
				JurusanID: jurusan.ID.String(),
				Name:      jurusan.Name,
				UpdatedAt: jurusan.UpdatedAt.Format(time.RFC3339),
			},
		})
	}
}
