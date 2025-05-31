package getbyid

import (
	"monitoring-guru/entities"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GetRuanganByID godoc
// @Summary Get ruangan by ID
// @Description Get a ruangan by its ID
// @Tags ruangan
// @Accept json
// @Produce json
// @Param id path string true "Ruangan ID"
// @Success 200 {object} GetRuanganByIDResponseWrapper
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/ruangan/{id} [get]
func GetRuanganByID(db *gorm.DB) fiber.Handler {
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

		return c.JSON(GetRuanganByIDResponseWrapper{
			Code:    200,
			Message: "Successfully fetched ruangan",
			Data: RuanganResponse{
				RuanganID: ruangan.ID.String(),
				Name:      ruangan.Name,
				CreatedAt: ruangan.CreatedAt.Format(time.RFC3339),
				UpdatedAt: ruangan.UpdatedAt.Format(time.RFC3339),
			},
		})
	}
}
