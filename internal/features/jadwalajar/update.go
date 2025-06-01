package jadwalajar

import (
	"fmt"
	e "monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// UpdateJadwalAjarRequestBody
// @Description Update jadwalajar request body
type UpdateJadwalAjarRequest struct {
	// @Description ID of the jadwalajar
	// @Required true
	// @Example "123456789"
	ID string `json:"id"`
	// @Description Guru ID of the jadwalajar
	// @Required true
	// @Example "123456789"
	GuruID string `json:"guru_id"`
	// @Description Mapel ID of the jadwalajar
	// @Required true
	// @Example "John Doe"
	MapelID string `json:"mapel_id"`
	// @Description Kelas ID of the jadwalajar
	// @Required true
	// @Example "adasd323"
	KelasID string `json:"kelas_id"`
	// @Description Hari of the jadwalajar
	// @Required true
	// @Example "Senin"
	Hari string `json:"hari"`
	// @Description Jam Mulai of the jadwalajar
	// @Required true
	// @Example "08:00"
	JamMulai string `json:"jam_mulai"`
	// @Description Jam Selesai of the jadwalajar
	// @Required true
	// @Example "10:00"
	JamSelesai string `json:"jam_selesai"`
	// @Description Last Editor of the jadwalajar
	// @Example "John Doe"
	LastEditor string `json:"last_editor"`
}


// UpdateJadwalAjarRequest godoc
// @summary Update Jadwalajar request body
// @Description	Update Jadwalajar baru request body
// @Tags			Jadwalajar
// @Accept			json
// @Produce		json
// @Param			request	body		UpdateJadwalAjarRequest	true	"Update jadwalajar request body"
// @Success		200		{object}	JadwalajarResponseWrapper
// @Failure		400		{object}	JadwalajarResponseWrapper
// @Failure		500		{object}	JadwalajarResponseWrapper
// @Router			/api/jadwalajar [put]
func (h *JadwalajarHandler) UpdateJadwalajar() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req UpdateJadwalAjarRequest
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

		jadwalID, err := parseUUID(req.ID, "ID")
		if err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, err.Error(), nil))
		}

		guruID, err := parseUUID(req.GuruID, "GuruID")
		if err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, err.Error(), nil))
		}

		mapelID, err := parseUUID(req.MapelID, "MapelID")
		if err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, err.Error(), nil))
		}

		kelasID, err := parseUUID(req.KelasID, "KelasID")
		if err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, err.Error(), nil))
		}

		jadwalajar := e.JadwalAjar{
			ID:         jadwalID,
			GuruID:     guruID,
			MapelID:    mapelID,
			KelasID:    kelasID,
			Hari:       req.Hari,
			JamMulai:   req.JamMulai,
			JamSelesai: req.JamSelesai,
			LastEditor: req.LastEditor,
		}

		if err := h.Service.UpdateJadwalajar(&jadwalajar); err != nil {
			return c.Status(500).JSON(e.ErrorResponse[any](500, err.Error(), nil))
		}

		return c.JSON(e.SuccessResponse(&jadwalajar))
	}
}