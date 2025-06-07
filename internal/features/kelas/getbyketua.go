package kelas

import (
	e "monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
)

// GetKelasByKetuaOrWakilHandler godoc
// @Summary Get kelas by ketua or wakil
// @Description Get kelas by ketua or wakil
// @Tags Kelas
// @Security BearerAuth
// @Produce json
// @Param ketua_kelas_id query string true "Ketua Kelas ID"
// @Success 200 {object} KelasResponseWrapper
// @Failure 404 {object} KelasResponseWrapper
// @Router /api/kelas/ketua [get]
func (h *KelasHandler) GetKelasByKetuaOrWakilHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ketuaKelasID := c.Query("ketua_kelas_id")

		if ketuaKelasID == "" {
			return c.Status(400).JSON(e.ErrorResponse[any](400, "Ketua Kelas ID is required", nil))
		}

		kelas, err := h.Service.GetKelasByKetuaOrWakil(ketuaKelasID)
		if err != nil {
			return c.Status(404).JSON(e.ErrorResponse[any](404, "Kelas not found", nil))
		}

		return c.JSON(e.SuccessResponse(kelas))
	}
}
