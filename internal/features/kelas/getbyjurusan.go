package kelas

import (
	e "monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
)

// GetKelasByJurusanHandler godoc
// @Summary Get kelas by jurusan
// @Description Get kelas by jurusan
// @Tags kelas
// @Security BearerAuth
// @Produce json
// @Param jurusan_id query string true "Jurusan ID"
// @Success 200 {object} []KelasResponse
// @Failure 500 {object} KelasResponseWrapper
// @Router /api/kelas/jurusan [get]
func (h *KelasHandler) GetKelasByJurusanHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		jurusanID := c.Query("jurusan_id")

		if jurusanID == "" {
			return c.Status(400).JSON(e.ErrorResponse[any](400, "Jurusan ID is required", nil))
		}

		kelasList, err := h.Service.GetKelasByJurusan(jurusanID)
		if err != nil {
			return c.Status(500).JSON(e.ErrorResponse[any](500, "Internal server error", nil))
		}

		return c.JSON(e.SuccessResponse(&kelasList))
	}
}
