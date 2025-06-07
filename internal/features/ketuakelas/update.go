package ketuakelas

import (
	e "monitoring-guru/entities"
	"monitoring-guru/infrastructure/repositories/user"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UpdateKetuaKelasRequest struct {
	ID       string `json:"id"`
	Name     string `json:"nama,omitempty"`
	Nisn     string `json:"nisn,omitempty"`
	Password string `json:"password,omitempty"`
}

// UpdateKetuaKelasHandler godoc
// @summary Update Ketua Kelas
// @description Update Ketua Kelas by ID, only fields provided will be updated
// @tags Ketua Kelas
// @security BearerAuth
// @Accept json
// @Produce json
// @Param request body UpdateKetuaKelasRequest true "Ketua Kelas update data"
// @Success 200 {object} KetuaKelasResponseWrapper
// @Failure 400 {object} KetuaKelasResponseWrapper
// @Failure 404 {object} KetuaKelasResponseWrapper
// @Failure 500 {object} KetuaKelasResponseWrapper
// @Router /api/ketua-kelas [put]
func (h *KetuaKelasHandler) UpdateKetuaKelasHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var req UpdateKetuaKelasRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(e.ErrorResponse[any](400, "Invalid request body", nil))
		}

		if req.ID == "" {
			return c.Status(fiber.StatusBadRequest).JSON(e.ErrorResponse[any](400, "ID is required", nil))
		}

		uid, err := uuid.Parse(req.ID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(e.ErrorResponse[any](400, "Invalid ID format", nil))
		}

		existing, err := h.Service.GetKetuaKelas(uid.String())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(e.ErrorResponse[any](500, "Failed to get ketua kelas", nil))
		}

		if existing == nil {
			return c.Status(fiber.StatusNotFound).JSON(e.ErrorResponse[any](404, "Ketua kelas not found", nil))
		}

		if req.Name != "" { existing.Name = req.Name }
		if req.Nisn != "" { existing.Nisn = req.Nisn }
		if req.Password != "" {
			if hashed, _ := user.HashPassword(req.Password); hashed != "" {
				existing.Password = hashed
			}
		}

		if err := h.Service.UpdateKetuaKelas(existing); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(e.ErrorResponse[any](500, err.Error(), nil))
		}

		return c.JSON(e.SuccessResponse(h.Service.ResponseKetuaKelasMapper(existing)))
	}
}
