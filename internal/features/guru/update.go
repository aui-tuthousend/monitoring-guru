package guru

import (
	e "monitoring-guru/entities"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UpdateGuruRequest struct {
	ID 		string `json:"id"`
	Name     string `json:"name,omitempty"`
	Nip      string `json:"nip,omitempty"`
	Jabatan  string `json:"jabatan,omitempty"`
}

// UpdateGuruHandler godoc
// @Summary Update guru data
// @Description Update a guru by ID, only fields provided will be updated
// @Tags Guru
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body UpdateGuruRequest true "Guru update data"
// @Success 200 {object} GuruResponseWrapper
// @Failure 400 {object} GuruResponseWrapper
// @Failure 500 {object} GuruResponseWrapper
// @Router /api/guru [put]
func (h *GuruHandler) UpdateGuruHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var req UpdateGuruRequest
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

		existing, err := h.Service.GetGuru(uid.String())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(e.ErrorResponse[any](500, "Guru not found", nil))
		}

		if strings.TrimSpace(req.Nip) == "" || strings.TrimSpace(req.Name) == "" || (req.Jabatan != "guru" && req.Jabatan != "kepala_sekolah") {
			return c.Status(400).JSON(e.ErrorResponse[any](400, "Invalid input data", nil))
		}

		existing.Name = req.Name
		existing.Nip = req.Nip
		existing.Jabatan = req.Jabatan

		if err := h.Service.UpdateGuru(existing); err != nil {
			return c.Status(500).JSON(e.ErrorResponse[any](500, err.Error(), nil))
		}

		return c.JSON(e.SuccessResponse(h.Service.ResponseGuruMapper(existing)))
	}
}
