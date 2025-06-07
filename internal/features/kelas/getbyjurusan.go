package kelas

import (
	e "monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetKelasByJurusanHandler godoc
// @Summary Get kelas by jurusan
// @Description Get kelas by jurusan
// @Tags Kelas
// @Security BearerAuth
// @Produce json
// @Param id path string true "Jurusan ID"
// @Success 200 {object} []KelasResponse
// @Failure 500 {object} KelasResponseWrapper
// @Router /api/kelas/jurusan/{id} [get]
func (h *KelasHandler) GetKelasByJurusanHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		parsedId, err := uuid.Parse(id)
		if err != nil {
			return c.Status(404).JSON(e.ErrorResponse[any](404, "Invalid Guid", nil))
		}

		kelasList, err := h.Service.GetKelasByJurusan(parsedId)
		if err != nil {
			return c.Status(500).JSON(e.ErrorResponse[any](500, "Internal server error", nil))
		}

		return c.JSON(e.SuccessResponse(&kelasList))
	}
}
