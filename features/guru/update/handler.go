package update

import (
	"monitoring-guru/infrastructure/repositories/guru"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// UpdateGuruHandler godoc
// @Summary Update guru data
// @Description Update a guru by ID, only fields provided will be updated
// @Tags guru
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Guru ID"
// @Param data body update.UpdateGuruRequest true "Guru update data"
// @Success 200 {object} update.UpdateGuruResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/guru/{id} [put]
func UpdateGuruHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		var req UpdateGuruRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		// Ambil data lama dari DB
		existing, err := guru.GetGuru(db, id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Guru not found",
			})
		}

		// Update field jika tidak nil
		if req.Name != nil {
			existing.Name = *req.Name
		}
		if req.Nip != nil {
			existing.Nip = *req.Nip
		}
		if req.Password != nil {
			existing.Password = *req.Password // ideally hash this if not already
		}
		if req.Jabatan != nil {
			existing.Jabatan = *req.Jabatan
		}

		if err := db.Save(&existing).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update guru",
			})
		}

		return c.JSON(UpdateGuruResponse{
			ID:        existing.ID.String(),
			Name:      existing.Name,
			Nip:       existing.Nip,
			Jabatan:   existing.Jabatan,
			CreatedAt: existing.CreatedAt,
			UpdatedAt: existing.UpdatedAt,
		})
	}
}
