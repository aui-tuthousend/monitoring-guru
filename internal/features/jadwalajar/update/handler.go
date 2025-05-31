package update

import (
	"fmt"
	e "monitoring-guru/entities"
	r "monitoring-guru/infrastructure/repositories/jadwalajar"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UpdateJadwalAjarRequest godoc
// @summary Update Jadwalajar request body
// @Description	Update Jadwalajar baru request body
// @Tags			Jadwalajar
// @Accept			json
// @Produce		json
// @Param			request	body		UpdateJadwalAjarRequest	true	"Update jadwalajar request body"
// @Success		200		{object}	UpdateJadwalAjarResponseWrapper
// @Failure		400		{object}	UpdateJadwalAjarResponseWrapper
// @Failure		500		{object}	UpdateJadwalAjarResponseWrapper
// @Router			/api/jadwalajar [put]
func UpdateJadwalAjar(db *gorm.DB) fiber.Handler {
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

		if err := r.UpdateJadwalajar(db, &jadwalajar); err != nil {
			return c.Status(500).JSON(e.ErrorResponse[any](500, err.Error(), nil))
		}

		return c.JSON(e.SuccessResponse(&jadwalajar))
	}
}
