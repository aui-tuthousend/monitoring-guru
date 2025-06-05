package absenmasuk

import (
	e "monitoring-guru/entities"
	"monitoring-guru/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CreateAbsenMasukRequest struct {
	GuruID     string `json:"guru_id"`
	JadwalAjarID string `json:"jadwal_ajar_id"`
	RuanganID    string `json:"ruangan_id"`
	Tanggal      string `json:"tanggal"`
	JamMasuk     string `json:"jam_masuk"`
}

func (h *AbsenMasukHandler) CreateAbsenMasuk() fiber.Handler {

	return func (c *fiber.Ctx) error  {
		var req CreateAbsenMasukRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, err.Error(), nil))
		}
		
		guruID, err := utils.ParseUUID(req.GuruID)
		if err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, err.Error(), nil))
		}
		jadwalajarID, err := utils.ParseUUID(req.JadwalAjarID)
		if err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, err.Error(), nil))
		}
		ruanganID, err := utils.ParseUUID(req.RuanganID)
		if err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, err.Error(), nil))
		}

		jamMasuk, err := utils.ParseJamString(req.JamMasuk)
		if err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, err.Error(), nil))
		}
		
		if err := h.Service.CreateAbsenMasuk(&e.AbsenMasuk{
			ID:         uuid.New(),
			GuruID:     guruID,
			JadwalAjarID: jadwalajarID,
			RuanganID:    ruanganID,
			Tanggal:      req.Tanggal,
			JamMasuk:     jamMasuk,
		}); err != nil {
			return c.Status(500).JSON(e.ErrorResponse[any](500, err.Error(), nil))
		}
		
		return c.JSON(e.SuccessResponse(&CreateAbsenMasukResponse{
			ID: req.ID,
		}))
		
	}

}