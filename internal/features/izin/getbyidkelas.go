package izin

import (
	e "monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *IzinHandler) GetAllIzinKelasHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("kelas_id")
		parsedId, err := uuid.Parse(id)
		if err != nil {
			return c.Status(500).JSON(e.ErrorResponse[any](500, "UUID Invalid", nil))
		}

		izins, err := h.Service.GetAllIzinKelas(parsedId)
		if err != nil {
			return c.Status(500).JSON(e.ErrorResponse[any](500, "Internal server error", nil))
		}

		return c.JSON(e.SuccessResponse(&izins))
	}
}