package ruangan

import (
	e "monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// DeleteRuangan godoc
// @Summary Delete ruangan
// @Description Hapus ruangan berdasarkan ID
// @Tags ruangan
// @Accept json
// @Produce json
// @Param id path string true "Ruangan ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/ruangan/{id} [delete]
func (h *RuanganHandler) DeleteRuangan() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		ruanganID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(e.ErrorResponse[any](400, "Format ID tidak valid", nil))
		}

		if err := h.Service.DeleteRuangan(ruanganID.String()); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(e.ErrorResponse[any](500, "Gagal menghapus ruangan", nil))
		}

		return c.JSON(fiber.Map{
			"code":    fiber.StatusOK,
			"message": "Ruangan berhasil dihapus",
		})
	}
}
