package create

import (
	"fmt"
	e "monitoring-guru/entities"
	r "monitoring-guru/infrastructure/repositories/jadwalajar"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateJadwalAjarRequest godoc
// @summary Create Jadwalajar request body
// @Description	Create Jadwalajar baru request body
// @Tags			Jadwalajar
// @Accept			json
// @Produce		json
// @Param			request	body		CreateJadwalAjarRequest	true	"Create jadwalajar request body"
// @Success		200		{object}	CreateJadwalAjarResponseWrapper
// @Failure		400		{object}	CreateJadwalAjarResponseWrapper
// @Failure		500		{object}	CreateJadwalAjarResponseWrapper
// @Router			/api/jadwalajar [post]
func CreateJadwalAjar(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req CreateJadwalAjarRequest
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
			ID:         uuid.New(),
			GuruID:     guruID,
			MapelID:    mapelID,
			KelasID:    kelasID,
			Hari:       req.Hari,
			JamMulai:   req.JamMulai,
			JamSelesai: req.JamSelesai,
			LastEditor: req.LastEditor,
		}

		if err := r.CreateJadwalajar(db, &jadwalajar); err != nil {
			return c.Status(500).JSON(e.ErrorResponse[any](500, err.Error(), nil))
		}

		return c.JSON(e.SuccessResponse(&CreateJadwalAjarResponse{
			ID:         jadwalajar.ID.String(),
			GuruID:     jadwalajar.GuruID.String(),
			MapelID:    jadwalajar.MapelID.String(),
			KelasID:    jadwalajar.KelasID.String(),
			Hari:       jadwalajar.Hari,
			JamMulai:   jadwalajar.JamMulai,
			JamSelesai: jadwalajar.JamSelesai,
			LastEditor: jadwalajar.LastEditor,
		}))
	}
}
