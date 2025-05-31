package update

import (
	"monitoring-guru/entities"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UpdateRuangan godoc
// @Summary Update ruangan
// @Description Update a ruangan by ID
// @Tags ruangan
// @Accept json
// @Produce json
// @Param id path string true "Ruangan ID"
// @Param request body UpdateRuanganRequest true "Ruangan body"
// @Success 200 {object} UpdateRuanganResponseWrapper
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/ruangan/{id} [put]
func UpdateRuangan(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		ruanganID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid ID format"})
		}

		var req UpdateRuanganRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}

		if req.Name == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Name is required"})
		}

		var ruangan entities.Ruangan
		if err := db.First(&ruangan, "id = ?", ruanganID).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Ruangan not found"})
		}

		ruangan.Name = req.Name
		ruangan.UpdatedAt = time.Now()

		if err := db.Save(&ruangan).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to update ruangan"})
		}

		return c.JSON(UpdateRuanganResponseWrapper{
			Code:    200,
			Message: "Ruangan updated successfully",
			Data: UpdateRuanganResponse{
				RuanganID: ruangan.ID.String(),
				Name:      ruangan.Name,
				UpdatedAt: ruangan.UpdatedAt.Format(time.RFC3339),
			},
		})
	}
}
