package statuskelas

import (
	"github.com/gofiber/fiber/v2"
	e "monitoring-guru/entities"
)
	

func (h *StatusKelasHandler) GetAllStatusKelas() fiber.Handler {
	return func(c *fiber.Ctx) error {

		statusKelas, err := h.Service.GetAllStatusKelas()
		if err != nil {
			return c.Status(500).JSON(e.ErrorResponse[any](500, "Internal server error", nil))
		}
		return c.JSON(e.SuccessResponse(&statusKelas))
	}
}