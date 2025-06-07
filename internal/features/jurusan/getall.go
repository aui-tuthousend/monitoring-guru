package jurusan

import (
	e "monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
)

// GetAllJurusan godoc
// @Summary Get all jurusan
// @Description Get list of all jurusan
// @Tags Jurusan
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} []JurusanResponse
// @Failure 500 {object} JurusanResponseWrapper
// @Router /api/jurusan [get]
func (h *JurusanHandler) GetAllJurusan() fiber.Handler {
	return func(c *fiber.Ctx) error {
		jurusanList, err := h.Service.GetAllJurusan()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(e.ErrorResponse[any](500, "Failed to fetch jurusan data", nil))
		}

		var responses []JurusanResponse
		for _, jurusan := range jurusanList {
			responses = append(responses, *h.Service.ResponseJurusanMapper(&jurusan))
		}

		return c.JSON(e.SuccessResponse(&responses))
	}
}