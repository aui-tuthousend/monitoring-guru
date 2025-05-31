package getbyid

import (
	"monitoring-guru/entities"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GetJurusanByID godoc
// @Summary Get jurusan by ID
// @Description Get jurusan details by ID
// @Tags jurusan
// @Accept json
// @Produce json
// @Param id path string true "Jurusan ID"
// @Success 200 {object} GetJurusanResponseWrapper
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/jurusan/{id} [get]
func GetJurusanByID(db *gorm.DB) fiber.Handler {
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

		return c.JSON(GetJurusanResponseWrapper{
			Code:    fiber.StatusOK,
			Message: "Successfully retrieved jurusan",
			Data: JurusanResponse{
				JurusanID: jurusan.ID.String(),
				Name:      jurusan.Name,
				CreatedAt: jurusan.CreatedAt.Format(time.RFC3339),
				UpdatedAt: jurusan.UpdatedAt.Format(time.RFC3339),
			},
		})
	}
}
