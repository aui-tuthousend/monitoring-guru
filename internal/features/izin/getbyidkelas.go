package izin

import (
	e "monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *IzinHandler) GetAllIzinKelasHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("kelas_id")
		izins, err := h.Service.GetAllIzinKelas(uuid.MustParse(id))
		if err != nil {
			return c.Status(500).JSON(e.ErrorResponse[any](500, "Internal server error", nil))
		}

		return c.JSON(e.SuccessResponse(&izins))
	}
}