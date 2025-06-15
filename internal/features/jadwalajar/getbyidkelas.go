package jadwalajar

import (
	e "monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetJadwalAjarKelasRequestBody
// @Description Get jadwalajar by ID Kelas request body
type GetJadwalAjarKelasRequest struct {
	// @Description Kelas ID of the jadwalajar
	// @Required true
	// @Example "123456789"
	KelasID string `json:"kelas_id"`
	// @Description Hari of the jadwalajar
	// @Required true
	// @Example "Senin"
	Hari string `json:"hari"`
}

// GetJadwalAjarGuruRequest godoc
// @summary Get Jadwalajar by ID Kelas request body
// @Description	Get Jadwalajar by ID Kelas request body
// @Tags			Jadwalajar
// @Security     BearerAuth
// @Accept			json
// @Produce		json
// @Param			id path		string	true	"Kelas ID"
// @Param			hari path	string	false	"Hari"
// @Success		200		{object}	[]JadwalajarResponseWrapper
// @Failure		400		{object}	JadwalajarResponseWrapper
// @Failure		500		{object}	JadwalajarResponseWrapper
// @Router			/api/jadwalajar/kelas/{id}/{hari} [get]
func (h *JadwalajarHandler) GetJadwalAjarByIDKelas() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		hari := c.Params("hari")

		if hari == "%7Bhari%7D" || hari == "undefined" || hari == "" {
			hari = ""
		}

		parsedId, err := uuid.Parse(id)
        if err != nil {
            return c.Status(400).JSON(e.ErrorResponse[any](400, err.Error(), nil))
        }

		jadwalajarList, err := h.Service.GetJadwalajarByIDKelas(parsedId, hari)
        if err != nil {
            return c.Status(500).JSON(e.ErrorResponse[any](500, err.Error(), nil))
        }

		return c.JSON(e.SuccessResponse(&jadwalajarList))
	}
}
