package create

import (
	"monitoring-guru/entities"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateRuangan godoc
// @Summary Create ruangan
// @Description Create a new ruangan
// @Tags ruangan
// @Accept json
// @Produce json
// @Param request body CreateRuanganRequest true "Ruangan body"
// @Success 200 {object} CreateRuanganResponseWrapper
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/ruangan [post]
func CreateRuangan(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req CreateRuanganRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}

		if req.Name == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Name is required"})
		}

		ruangan := entities.Ruangan{
			ID:        uuid.New(),
			Name:      req.Name,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := db.Create(&ruangan).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to create ruangan"})
		}

		return c.JSON(CreateRuanganResponseWrapper{
			Code:    200,
			Message: "Ruangan created successfully",
			Data: CreateRuanganResponse{
				RuanganID: ruangan.ID.String(),
				Name:      ruangan.Name,
			},
		})
	}
}
