package getall

import (
	e "monitoring-guru/entities"
	r "monitoring-guru/infrastructure/repositories/jadwalajar"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetAllJadwalAjar godoc
// @summary Get All Jadwalajar request body
// @Description	Get All Jadwalajar request body
// @Tags			Jadwalajar
// @Accept			json
// @Produce		json
// @Success		200		{object}	GetAllJadwalAjarResponseWrapper
// @Failure		400		{object}	GetAllJadwalAjarResponseWrapper
// @Router			/api/jadwalajar [get]
func GetAllJadwalAjar(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		jadwalajarList, err := r.GetAllJadwalajar(db)
        if err != nil {
            return c.Status(400).JSON(e.ErrorResponse[any](400, "Jadwal Ajar Not Found", nil))
        }

		return c.JSON(e.SuccessResponse(&jadwalajarList))
	}
}
