package jadwalajar

import (
	e "monitoring-guru/entities"
	"monitoring-guru/internal/features/guru"
	"monitoring-guru/internal/features/kelas"
	"monitoring-guru/internal/features/mapel"
	"monitoring-guru/internal/features/ruangan"
	"monitoring-guru/utils"

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
	// @Description Ruangan ID of the jadwalajar
	// @Required true
	// @Example "adasd323"
	RuanganID string `json:"ruangan_id"`
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
// @Security     BearerAuth
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

		jadwalID, _ := utils.ParseUUID(req.ID)
		guruID, _ := utils.ParseUUID(req.GuruID)
		mapelID, _ := utils.ParseUUID(req.MapelID)
		kelasID, _ := utils.ParseUUID(req.KelasID)
		ruanganID, _ := utils.ParseUUID(req.RuanganID)

		if jadwalID == uuid.Nil || guruID == uuid.Nil || mapelID == uuid.Nil || kelasID == uuid.Nil || ruanganID == uuid.Nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, "Invalid ID format", nil))
		}

		guruResponse, err := h.GuruService.GetGuru(guruID.String())
		if err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, "Guru not found", nil))
		}

		mapelResponse, err := h.MapelService.GetMapelByID(mapelID)
		if err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, "Mapel not found", nil))
		}

		kelasResponse, err := h.KelasService.GetKelasByID(kelasID)
		if err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, "Kelas not found", nil))
		}

		ruanganResponse, err := h.RuanganService.GetRuanganByID(ruanganID.String())
		if err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, "Ruangan not found", nil))
		}

		jamMulai, _ := utils.ParseJamString(req.JamMulai)
		jamSelesai, _ := utils.ParseJamString(req.JamSelesai)

		if jamMulai.IsZero() || jamSelesai.IsZero() {
			return c.Status(400).JSON(e.ErrorResponse[any](400, "Invalid time format", nil))
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

		return c.JSON(e.SuccessResponse(&JadwalajarResponse{
			ID:         jadwalajar.ID.String(),
			Guru:       &guru.GuruMiniResponse{
				ID:       guruResponse.ID.String(),
				Name:     guruResponse.Name,
			},
			Mapel:      &mapel.MapelMiniResponse{
				ID:       mapelResponse.ID,
				Name:     mapelResponse.Name,
			},
			Kelas:      &kelas.KelasMiniResponse{
				ID:       kelasResponse.ID,
				Name:     kelasResponse.Name,
			},
			Ruangan: &ruangan.RuanganResponse{
				ID:       ruanganResponse.ID.String(),
				Name:     ruanganResponse.Name,
			},
			Hari:       req.Hari,
			JamMulai:   req.JamMulai,
			JamSelesai: req.JamSelesai,
		}))
	}
}