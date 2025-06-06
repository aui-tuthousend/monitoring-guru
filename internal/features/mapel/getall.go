package mapel

import (
	e "monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
)

// GetAllMapelHandler godoc
// @Summary Get all mapel
// @Description Get all mata pelajaran
// @Tags Mapel
// @Security BearerAuth
// @Produce json
// @Success 200 {object} []MapelResponse
// @Failure 500 {object} MapelResponseWrapper
// @Router /api/mapel [get]
func (h *MapelHandler) GetAllMapel() fiber.Handler {
	return func(c *fiber.Ctx) error {
		mapelList, err := h.Service.GetAllMapel()
		if err != nil {
			return c.Status(500).JSON(e.ErrorResponse[any](500, "Internal server error", nil))
		}

		var responses []MapelResponse
		for _, mapel := range mapelList {
			responses = append(responses, *h.Service.ResponseMapelMapper(&mapel))
		}

		return c.JSON(e.SuccessResponse(&responses))
	}
}
