package getall

import (
	r "monitoring-guru/infrastructure/repositories/guru"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetAllGuruHandler godoc
// @Summary Get all guru
// @Description Get all guru
// @Tags guru
// @Security BearerAuth
// @Produce json
// @Success 200 {object} []GetAllGuruResponse
// @Failure 500 {object} map[string]string
// @Router /api/guru [get]
func GetAllGuruHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		gurus, err := r.GetAllGuru(db)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
		}

		// var guruResponses []GetAllGuruResponse
		// for _, guru := range gurus {
		// 	guruResponses = append(guruResponses, GetAllGuruResponse{
		// 		ID:   guru.ID,
		// 		NIP:  guru.Nip,
		// 		Name: guru.Name,
		// 	})
		// }

		return c.JSON(gurus)
	}
}
