package kelas

import (
	e "monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetKelasByIDHandler godoc
// @Summary Get kelas by ID
// @Description Get kelas by ID
// @Tags Kelas
// @Security BearerAuth
// @Produce json
// @Param id path string true "Kelas ID"
// @Success 200 {object} KelasResponseWrapper
// @Failure 404 {object} KelasResponseWrapper
// @Router /api/kelas/{id} [get]
func (h *KelasHandler) GetKelasByIDHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		parsedId, err := uuid.Parse(id)
		if err != nil {
			return c.Status(404).JSON(e.ErrorResponse[any](404, "Invalid Guid", nil))
		}

		kelas, err := h.Service.GetKelasByID(parsedId)
		if err != nil {
			return c.Status(404).JSON(e.ErrorResponse[any](404, "Kelas not found", nil))
		}

		return c.JSON(e.SuccessResponse(&kelas))
	}
}
