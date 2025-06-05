package jadwalajar

import (
	e "monitoring-guru/entities"
	"monitoring-guru/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// CreateJadwalAjarRequestBody
// @Description Create jadwalajar request body
type CreateJadwalAjarRequest struct {
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

// CreateJadwalAjarRequest godoc
// @summary Create Jadwalajar request body
// @Description	Create Jadwalajar baru request body
// @Tags			Jadwalajar
// @Accept			json
// @Produce		json
// @Param			request	body		CreateJadwalAjarRequest	true	"Create jadwalajar request body"
// @Success		200		{object}	JadwalajarResponseWrapper
// @Failure		400		{object}	JadwalajarResponseWrapper
// @Failure		500		{object}	JadwalajarResponseWrapper
// @Router			/api/jadwalajar [post]
func (h *JadwalajarHandler) CreateJadwalAjar() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req CreateJadwalAjarRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, err.Error(), nil))
		}


		guruID, err := utils.ParseUUID(req.GuruID)
		if err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, err.Error(), nil))
		}

		mapelID, err := utils.ParseUUID(req.MapelID)
		if err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, err.Error(), nil))
		}

		kelasID, err := utils.ParseUUID(req.KelasID)
		if err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, err.Error(), nil))
		}

		jadwalajar := e.JadwalAjar{
			ID:         uuid.New(),
			GuruID:     guruID,
			MapelID:    mapelID,
			KelasID:    kelasID,
			Hari:       req.Hari,
			JamMulai:   req.JamMulai,
			JamSelesai: req.JamSelesai,
			LastEditor: req.LastEditor,
		}

		if err := h.Service.CreateJadwalajar(&jadwalajar); err != nil {
			return c.Status(500).JSON(e.ErrorResponse[any](500, err.Error(), nil))
		}

		return c.JSON(e.SuccessResponse(&JadwalajarResponse{
			ID:         jadwalajar.ID.String(),
			// GuruID:     jadwalajar.GuruID.String(),
			// MapelID:    jadwalajar.MapelID.String(),
			// KelasID:    jadwalajar.KelasID.String(),
			Hari:       jadwalajar.Hari,
			JamMulai:   jadwalajar.JamMulai,
			JamSelesai: jadwalajar.JamSelesai,
			LastEditor: jadwalajar.LastEditor,
		}))
	}
}
