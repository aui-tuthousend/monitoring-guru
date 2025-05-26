package getall

import (
	"monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetAllRuangan godoc
// @Summary Get all ruangan
// @Description Get all ruangan records
// @Tags ruangan
// @Accept json
// @Produce json
// @Success 200 {object} GetAllRuanganResponseWrapper
// @Failure 500 {object} map[string]string
// @Router /api/ruangan [get]
func GetAllRuangan(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var ruanganList []entities.Ruangan
		if err := db.Find(&ruanganList).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch ruangan"})
		}

		response := make([]RuanganResponse, len(ruanganList))
		for i, ruangan := range ruanganList {
			response[i] = RuanganResponse{
				RuanganID: ruangan.ID.String(),
				Name:      ruangan.Name,
			}
		}

		return c.JSON(GetAllRuanganResponseWrapper{
			Code:    200,
			Message: "Successfully fetched all ruangan",
			Data:    response,
		})
	}
}
