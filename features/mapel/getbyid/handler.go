package getbyid

import (
	"monitoring-guru/entities"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GetMapelByID godoc
// @Summary Get mapel by ID
// @Description Get a mata pelajaran by its ID
// @Tags mapel
// @Accept json
// @Produce json
// @Param id path string true "Mapel ID"
// @Success 200 {object} GetMapelByIDResponseWrapper
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/mapel/{id} [get]
func GetMapelByID(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		mapelID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
		}

		var mapel entities.Mapel
		if err := db.First(&mapel, "id = ?", mapelID).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Mapel not found"})
		}

		return c.JSON(GetMapelByIDResponseWrapper{
			Code:    fiber.StatusOK,
			Message: "Successfully fetched mapel",
			Data: MapelResponse{
				MapelID:   mapel.ID.String(),
				JurusanID: mapel.JurusanID.String(),
				Name:      mapel.Name,
				CreatedAt: mapel.CreatedAt.Format(time.RFC3339),
				UpdatedAt: mapel.UpdatedAt.Format(time.RFC3339),
			},
		})
	}
}
