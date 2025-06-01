package ketuakelas

import (
	"github.com/gofiber/fiber/v2"
)

// DeleteKetuaHandler godoc
// @summary Delete Ketua Kelas
// @Description	Delete Ketua Kelas by ID
// @Tags			ketua kelas
// @security BearerAuth
// produce json
// @Param			id	path		string	true	"Ketua Kelas ID"
// @Success		200	{object}	map[string]string
// @Failure		400	{object}	map[string]string
// @Failure		500	{object}	map[string]string
// @Router			/api/ketua-kelas/{id} [delete]
func (h *KetuaKelasHandler) DeleteKetuaHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == "" {
			return c.Status(400).JSON(fiber.Map{"error": "ID tidak boleh kosong"})
		}

		if err := h.Service.DeleteKetuaKelas(id); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal menghapus ketua kelas"})
		}

		return c.JSON(fiber.Map{"message": "Berhasil menghapus ketua kelas"})
	}
}
