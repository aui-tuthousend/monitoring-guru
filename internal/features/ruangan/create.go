package ruangan

import (
	e "monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// CreateRuanganRequest
// @Description Create ruangan request body
type CreateRuanganRequest struct {
	// @Description Nama ruangan
	// @Required true
	// @Example "Ruang 101"
	Name string `json:"name"`
}

// CreateRuangan godoc
// @summary     Create Ruangan request body
// @Description Buat ruangan baru
// @Tags        ruangan
// @Accept      json
// @Produce     json
// @Param       request body CreateRuanganRequest true "Create ruangan request body"
// @Success     200 {object} CreateRuanganResponseWrapper
// @Failure     400 {object} map[string]string
// @Failure     500 {object} map[string]string
// @Router      /api/ruangan [post]
func (h *RuanganHandler) CreateRuangan() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req CreateRuanganRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, err.Error(), nil))
		}

		if req.Name == "" {
			return c.Status(400).JSON(e.ErrorResponse[any](400, "Nama ruangan tidak boleh kosong", nil))
		}

		ruangan := e.Ruangan{
			ID:   uuid.New(),
			Name: req.Name,
		}

		if err := h.Service.CreateRuangan(&ruangan); err != nil {
			return c.Status(500).JSON(e.ErrorResponse[any](500, err.Error(), nil))
		}

		res := RuanganResponse{
			RuanganID: ruangan.ID.String(),
			Name:      ruangan.Name,
		}

		return c.JSON(e.SuccessResponse(&res))
	}
}
