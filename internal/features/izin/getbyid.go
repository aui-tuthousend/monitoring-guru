package izin

import (
	e "monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetIzinByID godoc
// @Summary Get izin by ID
// @Description Get izin details by ID
// @Tags Izin
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Izin ID"
// @Success 200 {object} IzinResponseWrapper
// @Failure 400 {object} IzinResponseWrapper
// @Failure 500 {object} IzinResponseWrapper
// @Router /api/izin/{id} [get]
func (h *IzinHandler) GetIzinByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		izinID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(e.ErrorResponse[any](400, "Invalid ID format", nil))
		}

		izin, err := h.Service.GetIzin(izinID.String())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(e.ErrorResponse[any](500, "Izin not found", nil))
		}

		return c.JSON(e.SuccessResponse(h.Service.ResponseIzinMapper(izin)))
	}
}
