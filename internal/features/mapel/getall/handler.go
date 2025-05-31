package getall

import (
	"monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetAllMapel godoc
// @Summary Get all mapel
// @Description Get all mata pelajaran
// @Tags mapel
// @Accept json
// @Produce json
// @Success 200 {object} GetAllMapelResponseWrapper
// @Failure 500 {object} map[string]string
// @Router /api/mapel [get]
func GetAllMapel(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var mapelList []entities.Mapel
		if err := db.Find(&mapelList).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch mapel"})
		}

		response := make([]MapelResponse, len(mapelList))
		for i, mapel := range mapelList {
			response[i] = MapelResponse{
				MapelID:   mapel.ID.String(),
				JurusanID: mapel.JurusanID.String(),
				Name:      mapel.Name,
			}
		}

		return c.JSON(GetAllMapelResponseWrapper{
			Code:    fiber.StatusOK,
			Message: "Successfully fetched all mapel",
			Data:    response,
		})
	}
}
