package update

import (
	"monitoring-guru/entities"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UpdateMapel godoc
// @Summary Update mapel
// @Description Update a mata pelajaran by ID
// @Tags mapel
// @Accept json
// @Produce json
// @Param id path string true "Mapel ID"
// @Param request body UpdateMapelRequest true "Mapel body"
// @Success 200 {object} UpdateMapelResponseWrapper
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/mapel/{id} [put]
func UpdateMapel(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		mapelID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
		}

		var req UpdateMapelRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
		}

		jurusanID, err := uuid.Parse(req.JurusanID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Jurusan ID format"})
		}

		if req.Name == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Name is required"})
		}

		var mapel entities.Mapel
		if err := db.First(&mapel, "id = ?", mapelID).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Mapel not found"})
		}

		mapel.JurusanID = jurusanID
		mapel.Name = req.Name
		mapel.UpdatedAt = time.Now()

		if err := db.Save(&mapel).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update mapel"})
		}

		return c.JSON(UpdateMapelResponseWrapper{
			Code:    fiber.StatusOK,
			Message: "Mapel updated successfully",
			Data: UpdateMapelResponse{
				MapelID:   mapel.ID.String(),
				JurusanID: mapel.JurusanID.String(),
				Name:      mapel.Name,
				UpdatedAt: mapel.UpdatedAt.Format(time.RFC3339),
			},
		})
	}
}
