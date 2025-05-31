package delete

import (
	"monitoring-guru/infrastructure/repositories/ketua"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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
func DeleteKetuaHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == "" {
			return c.Status(400).JSON(fiber.Map{"error": "ID tidak boleh kosong"})
		}

		if err := ketua.DeleteKetuaKelas(db, id); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Gagal menghapus ketua kelas"})
		}

		return c.JSON(fiber.Map{"message": "Berhasil menghapus ketua kelas"})
	}
}
