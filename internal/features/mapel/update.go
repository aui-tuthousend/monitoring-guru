package mapel

import (
	e "monitoring-guru/entities"
	"monitoring-guru/utils"

	"github.com/gofiber/fiber/v2"
)

// UpdateMapelRequest
// @Description Update mapel request body
type UpdateMapelRequest struct {
	// @Description ID of the mapel
	// @Required true
	// @Example "123e4567-e89b-12d3-a456-426614174000"
	ID string `json:"id"`
	// @Description Name of the mapel
	// @Required true
	// @Example "Fisika"
	Name string `json:"name"`
	// @Description Jurusan ID of the mapel
	// @Required true
	// @Example "123e4567-e89b-12d3-a456-426614174001"
	JurusanID string `json:"jurusan_id"`
}

// UpdateMapelHandler godoc
// @Summary Update mapel data
// @Description Update a mapel by ID
// @Tags Mapel
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body UpdateMapelRequest true "Update mapel request body"
// @Success 200 {object} MapelResponseWrapper
// @Failure 400 {object} MapelResponseWrapper
// @Failure 500 {object} MapelResponseWrapper
// @Router /api/mapel [put]
func (h *MapelHandler) UpdateMapel() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req UpdateMapelRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, err.Error(), nil))
		}

		mapelID, err := utils.ParseUUID(req.ID)
		if err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, err.Error(), nil))
		}
		jurusanID, err := utils.ParseUUID(req.JurusanID)
		if err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, err.Error(), nil))
		}

		mapel := e.Mapel{
			ID:        mapelID,
			Name:      req.Name,
			JurusanID: jurusanID,
		}

		if err := h.Service.UpdateMapel(&mapel); err != nil {
			return c.Status(500).JSON(e.ErrorResponse[any](500, err.Error(), nil))
		}

		return c.JSON(e.SuccessResponse(&mapel))
	}
}
