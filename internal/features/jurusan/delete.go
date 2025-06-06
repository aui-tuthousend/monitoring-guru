package jurusan

import (
	e "monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// DeleteJurusan godoc
// @Summary Delete jurusan
// @Description Delete jurusan by ID
// @Tags jurusan
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Jurusan ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/jurusan/{id} [delete]
func (h *JurusanHandler) DeleteJurusan() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		jurusanID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(e.ErrorResponse[any](400, "Invalid ID format", nil))
		}

		if err := h.Service.DeleteJurusan(jurusanID.String()); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(e.ErrorResponse[any](500, "Failed to delete jurusan", nil))
		}

		return c.JSON(fiber.Map{
			"code":    fiber.StatusOK,
			"message": "Jurusan deleted successfully",
		})
	}
}
