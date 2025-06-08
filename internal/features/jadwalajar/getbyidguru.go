package jadwalajar

import (
	e "monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetJadwalAjarGuruRequestBody
// @Description Get jadwalajar by ID Guru request body
type GetJadwalAjarGuruRequest struct {
	// @Description Guru ID of the jadwalajar
	// @Required true
	// @Example "123456789"
	GuruID string `json:"guru_id"`
	// @Description Hari of the jadwalajar
	// @Required true
	// @Example "Senin"
	Hari string `json:"hari"`
}

// GetJadwalAjarGuruRequest godoc
// @summary Get Jadwalajar by ID Guru request body
// @Description	Get Jadwalajar by ID Guru request body
// @Tags			Jadwalajar
// @Security     BearerAuth
// @Accept			json
// @Produce		json
// @Param			id path		string	true	"Guru ID"
// @Param			hari path	string	false	"Hari"
// @Success		200		{object}	[]JadwalajarResponse
// @Failure		400		{object}	JadwalajarResponseWrapper
// @Failure		500		{object}	JadwalajarResponseWrapper
// @Router			/api/jadwalajar/guru/{id}/{hari} [get]
func (h *JadwalajarHandler) GetJadwalAjarByIDGuru() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		hari := c.Params("hari")

		if hari == "%7Bhari%7D" {
			hari = ""
		}
		
		parsedId, err := uuid.Parse(id)
        if err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, err.Error(), nil))
        }

		jadwalajarList, err := h.Service.GetJadwalajarByIDGuru(parsedId, hari)
        if err != nil {
			return c.Status(500).JSON(e.ErrorResponse[any](500, err.Error(), nil))
        }

		return c.JSON(e.SuccessResponse(&jadwalajarList))
	}
}