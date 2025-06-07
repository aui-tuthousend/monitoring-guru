package ketuakelas

import (
	"github.com/gofiber/fiber/v2"
	e "monitoring-guru/entities"
)

// DeleteKetuaHandler godoc
// @summary Delete Ketua Kelas
// @Description	Delete Ketua Kelas by ID
// @Tags			Ketua Kelas
// @security BearerAuth
// @Accept json
// @Produce json
// @Param			id	path		string	true	"Ketua Kelas ID"
// @Success		200	{object}	map[string]string
// @Failure		400	{object}	map[string]string
// @Failure		500	{object}	map[string]string
// @Router			/api/ketua-kelas/{id} [delete]
func (h *KetuaKelasHandler) DeleteKetuaHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		if err := h.Service.DeleteKetuaKelas(id); err != nil {
			return c.Status(500).JSON(e.ErrorResponse[any](500, "Internal server error", nil))
		}

		return c.JSON(fiber.Map{"message": "Berhasil menghapus ketua kelas"})
	}
}
