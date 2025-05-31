package getall

import (
	"monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetAllJurusan godoc
// @Summary Get all jurusan
// @Description Get list of all jurusan
// @Tags jurusan
// @Accept json
// @Produce json
// @Success 200 {object} GetAllJurusanResponseWrapper
// @Failure 500 {object} map[string]string
// @Router /api/jurusan [get]
func GetAllJurusan(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var jurusans []entities.Jurusan
		if err := db.Find(&jurusans).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to fetch jurusan data",
			})
		}

		response := make([]JurusanResponse, len(jurusans))
		for i, jurusan := range jurusans {
			response[i] = JurusanResponse{
				JurusanID: jurusan.ID.String(),
				Name:      jurusan.Name,
			}
		}

		return c.JSON(GetAllJurusanResponseWrapper{
			Code:    fiber.StatusOK,
			Message: "Successfully retrieved all jurusan",
			Data:    response,
		})
	}
}
