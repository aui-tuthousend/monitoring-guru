package getbyidkelas

import (
	e "monitoring-guru/entities"
	r "monitoring-guru/infrastructure/repositories/jadwalajar"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GetJadwalAjarGuruRequest godoc
// @summary Get Jadwalajar by ID Kelas request body
// @Description	Get Jadwalajar by ID Kelas request body
// @Tags			Jadwalajar
// @Accept			json
// @Produce		json
// @Param			request	body		GetJadwalAjarKelasRequest	true	"Get jadwalajar kelas request body"
// @Success		200		{object}	GetByIDKelasResponseWrapper
// @Failure		400		{object}	GetByIDKelasResponseWrapper
// @Failure		500		{object}	GetByIDKelasResponseWrapper
// @Router			/api/jadwalajar/kelas [get]
func GetJadwalAjarByIDKelas(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req GetJadwalAjarKelasRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, err.Error(), nil))
		}
		
		id, err := uuid.Parse(req.KelasID)
        if err != nil {
            return c.Status(400).JSON(e.ErrorResponse[any](400, err.Error(), nil))
        }

		jadwalajarList, err := r.GetJadwalajarByIDKelas(db, id.String(), req.Hari)
        if err != nil {
            return c.Status(500).JSON(e.ErrorResponse[any](500, err.Error(), nil))
        }

		return c.JSON(e.SuccessResponse(&jadwalajarList))
	}
}
