package create

import (
	"monitoring-guru/entities"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateMapel godoc
// @Summary Create mapel
// @Description Create a new mata pelajaran
// @Tags mapel
// @Accept json
// @Produce json
// @Param request body CreateMapelRequest true "Mapel body"
// @Success 200 {object} CreateMapelResponseWrapper
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/mapel [post]
func CreateMapel(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req CreateMapelRequest
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

		mapel := entities.Mapel{
			ID:        uuid.New(),
			JurusanID: jurusanID,
			Name:      req.Name,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := db.Create(&mapel).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create mapel"})
		}

		return c.JSON(CreateMapelResponseWrapper{
			Code:    fiber.StatusOK,
			Message: "Mapel created successfully",
			Data: CreateMapelResponse{
				MapelID:   mapel.ID.String(),
				JurusanID: mapel.JurusanID.String(),
				Name:      mapel.Name,
			},
		})
	}
}
