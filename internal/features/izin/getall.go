package izin

import (
	e "monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
)

// GetAllIzinHandler godoc
// @Summary Get all izin
// @Description Get all izin
// @Tags Izin
// @Security BearerAuth
// @Produce json
// @Success 200 {object} []IzinResponse
// @Failure 500 {object} IzinResponseWrapper
// @Router /api/izin [get]
func (h *IzinHandler) GetAllIzinHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		izins, err := h.Service.GetAllIzin()
		if err != nil {
			return c.Status(500).JSON(e.ErrorResponse[any](500, "Internal server error", nil))
		}

		var responses []IzinResponse
		for _, izin := range izins {
			responses = append(responses, *h.Service.ResponseIzinMapper(&izin))
		}

		return c.JSON(e.SuccessResponse(&responses))
	}
}
