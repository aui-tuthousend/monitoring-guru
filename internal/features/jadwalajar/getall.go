package jadwalajar

import (
	e "monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
)

// GetAllJadwalAjar godoc
// @summary Get All Jadwalajar request body
// @Description	Get All Jadwalajar request body
// @Tags			Jadwalajar
// @Accept			json
// @Produce		json
// @Success		200		{object}	[]JadwalajarResponse
// @Failure		400		{object}	JadwalajarResponseWrapper
// @Router			/api/jadwalajar [get]
func (h *JadwalajarHandler) GetAllJadwalAjar() fiber.Handler {
	return func(c *fiber.Ctx) error {
		jadwalajarList, err := h.Service.GetAllJadwalajar()
        if err != nil {
            return c.Status(400).JSON(e.ErrorResponse[any](400, "Jadwal Ajar Not Found", nil))
        }

		return c.JSON(e.SuccessResponse(&jadwalajarList))
	}
}
