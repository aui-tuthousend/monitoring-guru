package jurusan

import (
	e "monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UpdateJurusanRequest struct {
	ID   string `json:"id"`
	Name string `json:"nama" validate:"required"`
}

// UpdateJurusan godoc
// @Summary Update jurusan
// @Description Update jurusan data
// @Tags Jurusan
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body UpdateJurusanRequest true "Jurusan data"
// @Success 200 {object} JurusanResponseWrapper
// @Failure 400 {object} JurusanResponseWrapper
// @Failure 500 {object} JurusanResponseWrapper
// @Router /api/jurusan [put]
func (h *JurusanHandler) UpdateJurusan() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req UpdateJurusanRequest
		if err := c.BodyParser(&req); err != nil {
			return c.JSON(e.ErrorResponse[any](400, "Invalid request body", nil))
		}

		jursanId, err := uuid.Parse(req.ID)
		if err != nil {
			return c.JSON(e.ErrorResponse[any](400, "Invalid ID format", nil))
		}

		jurusan, err := h.Service.GetJurusanByID(jursanId.String())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(e.ErrorResponse[any](500, "Jurusan not found", nil))
		}

		if req.Name == "" {
			return c.JSON(e.ErrorResponse[any](400, "Nama jurusan is required", nil))
		}

		jurusan.Name = req.Name
		if err := h.Service.UpdateJurusan(jurusan); err != nil {
			return c.JSON(e.ErrorResponse[any](500, "Failed to update jurusan", nil))
		}

		return c.JSON(e.SuccessResponse(h.Service.ResponseJurusanMapper(jurusan)))
	}
}