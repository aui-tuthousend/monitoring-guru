package ruangan

import (
	"fmt"
	e "monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// UpdateRuanganRequest
// @Description Update ruangan request body
type UpdateRuanganRequest struct {
	// @Description ID ruangan
	// @Required true
	// @Example "123e4567-e89b-12d3-a456-426614174000"
	ID string `json:"id"`
	// @Description Nama ruangan
	// @Required true
	// @Example "Ruang 102"
	Name string `json:"name"`
}

// UpdateRuanganHandler godoc
// @Summary Update ruangan data
// @Description Update a ruangan by ID
// @Tags ruangan
// @Accept json
// @Produce json
// @Param request body UpdateRuanganRequest true "Update ruangan request body"
// @Success 200 {object} UpdateRuanganResponseWrapper
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/ruangan [put]
func (h *RuanganHandler) UpdateRuangan() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req UpdateRuanganRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, err.Error(), nil))
		}

		parseUUID := func(idStr, field string) (uuid.UUID, error) {
			uid, err := uuid.Parse(idStr)
			if err != nil {
				return uuid.Nil, fmt.Errorf("%s tidak valid: %w", field, err)
			}
			return uid, nil
		}

		ruanganID, err := parseUUID(req.ID, "ID")
		if err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, err.Error(), nil))
		}

		if req.Name == "" {
			return c.Status(400).JSON(e.ErrorResponse[any](400, "Nama ruangan tidak boleh kosong", nil))
		}

		ruangan := e.Ruangan{
			ID:   ruanganID,
			Name: req.Name,
		}

		if err := h.Service.UpdateRuangan(&ruangan); err != nil {
			return c.Status(500).JSON(e.ErrorResponse[any](500, err.Error(), nil))
		}

		return c.JSON(e.SuccessResponse(&ruangan))
	}
}
