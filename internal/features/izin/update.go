package izin

import (
	e "monitoring-guru/entities"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UpdateIzinRequest struct {
	ID          string `json:"id"`
	TanggalIzin string `json:"tanggal_izin"`
	Pesan       string `json:"pesan"`
}

// UpdateIzinHandler godoc
// @Summary Update izin
// @Description Update data izin
// @Tags Izin
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body UpdateIzinRequest true "Izin data"
// @Success 200 {object} IzinResponseWrapper
// @Failure 400 {object} IzinResponseWrapper
// @Failure 500 {object} IzinResponseWrapper
// @Router /api/izin [put]
func (h *IzinHandler) UpdateIzinHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req UpdateIzinRequest
		if err := c.BodyParser(&req); err != nil {
			return c.JSON(e.ErrorResponse[any](400, "Invalid request body", nil))
		}

		izinID, err := uuid.Parse(req.ID)
		if err != nil {
			return c.JSON(e.ErrorResponse[any](400, "Invalid ID format", nil))
		}

		izin, err := h.Service.GetIzin(izinID.String())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(e.ErrorResponse[any](500, "Izin not found", nil))
		}

		if req.TanggalIzin == "" {
			return c.JSON(e.ErrorResponse[any](400, "Tanggal izin is required", nil))
		}

		if req.Pesan == "" {
			return c.JSON(e.ErrorResponse[any](400, "Pesan is required", nil))
		}

		tanggal, err := time.Parse("2006-01-02", req.TanggalIzin)
		if err != nil {
			return c.JSON(e.ErrorResponse[any](400, "Invalid tanggal_izin format, harus YYYY-MM-DD", nil))
		}

		izin.TanggalIzin = tanggal
		izin.Pesan = req.Pesan
		izin.UpdatedAt = time.Now()

		if err := h.Service.UpdateIzin(izin); err != nil {
			return c.JSON(e.ErrorResponse[any](500, "Failed to update izin", nil))
		}

		return c.JSON(e.SuccessResponse(h.Service.ResponseIzinMapper(izin)))
	}
}
