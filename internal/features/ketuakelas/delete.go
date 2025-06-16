package ketuakelas

import (
	e "monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
		uid, err := uuid.Parse(id)
			if err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, "Invalid ID format", nil))
		}

		if err := h.Service.DeleteKetuaKelas(uid); err != nil {
			return c.Status(500).JSON(e.ErrorResponse[any](500, "Internal server error", nil))
		}

		return c.JSON(fiber.Map{"message": "Berhasil menghapus ketua kelas"})
	}
}
