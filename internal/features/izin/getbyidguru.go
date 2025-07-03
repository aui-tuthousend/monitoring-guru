package izin

import (
	e "monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
)

func (h *IzinHandler) GetAllIzinGuruHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("nip")
		izins, err := h.Service.GetAllIzinGuru(id)
		if err != nil {
			return c.Status(500).JSON(e.ErrorResponse[any](500, "Internal server error", nil))
		}

		return c.JSON(e.SuccessResponse(&izins))
	}
}