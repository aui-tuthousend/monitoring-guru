package jurusan


import (
	e "monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetJurusanByID godoc
// @Summary Get jurusan by ID
// @Description Get jurusan details by ID
// @Tags jurusan
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Jurusan ID"
// @Success 200 {object} JurusanResponseWrapper
// @Failure 400 {object} JurusanResponseWrapper
// @Failure 500 {object} JurusanResponseWrapper
// @Router /api/jurusan/{id} [get]
func (h *JurusanHandler) GetJurusanByID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		jurusanID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(e.ErrorResponse[any](400, "Invalid ID format", nil))
		}

		jurusan, err := h.Service.GetJurusanByID(jurusanID.String())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(e.ErrorResponse[any](500, "Jurusan not found", nil))
		}

		return c.JSON(e.SuccessResponse(h.Service.ResponseJurusanMapper(jurusan)))
	}
}
