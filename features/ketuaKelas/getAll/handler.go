package getall

import (
	r "monitoring-guru/infrastructure/repositories/ketua"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetAllKetuaKelasHandler godoc
// @Summary Get all ketua kelas
// @Description Get all ketua kelas
// @Tags ketua kelas
// @Security BearerAuth
// @Produce json
// @Success 200 {object} []GetAllKetuaKelasResponse
// @Failure 500 {object} map[string]string
// @Router /api/ketua-kelas [get]
func GetAllKetuaKelasHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ketuaKelasList, err := r.GetAllKetuaKelas(db)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
		}

		var ketuaKelasResponses []GetAllKetuaKelasResponse
		for _, ketua := range ketuaKelasList {
			ketuaKelasResponses = append(ketuaKelasResponses, GetAllKetuaKelasResponse{
				ID:   ketua.ID,
				NISN: ketua.Nisn,
				Name: ketua.Name,
			})
		}

		return c.JSON(ketuaKelasResponses)
	}
}
