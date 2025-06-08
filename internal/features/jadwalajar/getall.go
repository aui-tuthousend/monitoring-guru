package jadwalajar

import (
	e "monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
)

// GetAllJadwalAjar godoc
// @summary Get All Jadwalajar request body
// @Description	Get All Jadwalajar request body
// @Tags			Jadwalajar
// @Security     BearerAuth
// @Accept			json
// @Produce		json
// @Success		200		{object}	[]JadwalajarResponse
// @Failure		500		{object}	JadwalajarResponseWrapper
// @Router			/api/jadwalajar [get]
func (h *JadwalajarHandler) GetAllJadwalAjar() fiber.Handler {
	return func(c *fiber.Ctx) error {
		jadwalajarList, err := h.Service.GetAllJadwalajar()
        if err != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(e.ErrorResponse[any](500, "Jadwal Ajar Not Found", nil))
        }

		return c.JSON(e.SuccessResponse(&jadwalajarList))
	}
}
