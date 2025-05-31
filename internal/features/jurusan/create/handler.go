package create

import (
	"monitoring-guru/entities"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateJurusan godoc
// @Summary Create jurusan
// @Description Create a new jurusan
// @Tags jurusan
// @Accept json
// @Produce json
// @Param request body CreateJurusanRequest true "Jurusan body"
// @Success 201 {object} CreateJurusanResponseWrapper
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/jurusan [post]
func CreateJurusan(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req CreateJurusanRequest
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

		jurusan := entities.Jurusan{
			ID:        uuid.New(),
			Name:      req.Name,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := db.Create(&jurusan).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create jurusan",
			})
		}

		return c.Status(fiber.StatusCreated).JSON(CreateJurusanResponseWrapper{
			Code:    fiber.StatusCreated,
			Message: "Jurusan created successfully",
			Data: CreateJurusanResponse{
				JurusanID: jurusan.ID.String(),
				Name:      jurusan.Name,
			},
		})
	}
}
