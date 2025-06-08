package jadwalajar

import (
	e "monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)


// GetJadwalajarByID godoc
// @Summary Get jadwalajar by ID
// @Description Get a jadwalajar by its ID
// @Tags Jadwalajar
// @Security     BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Jadwalajar ID"
// @Success 200 {object} JadwalajarResponseWrapper
// @Failure 400 {object} JadwalajarResponseWrapper
// @Failure 500 {object} JadwalajarResponseWrapper
// @Router /api/jadwalajar/{id} [get]
func (h *JadwalajarHandler) GetJadwalajarByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		jadwalajarID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(e.ErrorResponse[any](400, "Invalid GUID format", nil))
		}

		jadwalajar, err := h.Service.GetJadwalajarByID(jadwalajarID.String())
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(e.ErrorResponse[any](500, "Jadwalajar not found", nil))
		}

		return c.JSON(e.SuccessResponse(&jadwalajar))
	}
}
