package getbyidguru

import (
	e "monitoring-guru/entities"
	r "monitoring-guru/infrastructure/repositories/jadwalajar"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GetJadwalAjarGuruRequest godoc
// @summary Get Jadwalajar by ID Guru request body
// @Description	Get Jadwalajar by ID Guru request body
// @Tags			Jadwalajar
// @Accept			json
// @Produce		json
// @Param			request	body		GetJadwalAjarGuruRequest	true	"Get jadwalajar guru request body"
// @Success		200		{object}	GetByIDGuruResponseWrapper
// @Failure		400		{object}	GetByIDGuruResponseWrapper
// @Failure		500		{object}	GetByIDGuruResponseWrapper
// @Router			/api/jadwalajar/guru [get]
func GetJadwalAjarByIDGuru(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req GetJadwalAjarGuruRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, err.Error(), nil))
		}
		
		id, err := uuid.Parse(req.GuruID)
        if err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, err.Error(), nil))
        }

		jadwalajarList, err := r.GetJadwalajarByIDGuru(db, id.String(), req.Hari)
        if err != nil {
			return c.Status(500).JSON(e.ErrorResponse[any](500, err.Error(), nil))
        }

		return c.JSON(e.SuccessResponse(&jadwalajarList))
	}
}
