package mapel

import (
	"fmt"
	e "monitoring-guru/entities"
	"monitoring-guru/internal/features/jurusan"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// CreateMapelRequest
// @Description Create mapel request body
type CreateMapelRequest struct {
	// @Description Nama mapel
	// @Required true
	// @Example "Matematika"
	Name string `json:"name"`
	// @Description Jurusan ID dari mapel
	// @Required true
	// @Example "123e4567-e89b-12d3-a456-426614174000"
	JurusanID string `json:"jurusan_id"`
}

// CreateMapel godoc
// @summary     Create Mapel request body
// @Description Buat mapel baru
// @Tags        mapel
// @Accept      json
// @Produce     json
// @Param       request body CreateMapelRequest true "Create mapel request body"
// @Success     200     {object} MapelResponseWrapper
// @Failure     400     {object} map[string]string
// @Failure     500     {object} map[string]string
// @Router      /api/mapel [post]
func (h *MapelHandler) CreateMapel() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req CreateMapelRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, err.Error(), nil))
		}

		if req.Name == "" {
			return c.Status(400).JSON(e.ErrorResponse[any](400, "Nama mapel tidak boleh kosong", nil))
		}

		parseUUID := func(idStr string) (uuid.UUID, error) {
			uid, err := uuid.Parse(idStr)
			if err != nil {
				return uuid.Nil, fmt.Errorf("UUID %s tidak valid: %w", idStr, err)
			}
			return uid, nil
		}

		jurusanID, err := parseUUID(req.JurusanID)
		if err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, err.Error(), nil))
		}

		mapel := e.Mapel{
			ID:        uuid.New(),
			Name:      req.Name,
			JurusanID: jurusanID,
		}

		if err := h.Service.CreateMapel(&mapel); err != nil {
			return c.Status(500).JSON(e.ErrorResponse[any](500, err.Error(), nil))
		}

		var created e.Mapel
		if err := h.Service.DB.
			Preload("Jurusan").
			First(&created, "id = ?", mapel.ID).
			Error; err != nil {
			return c.Status(500).JSON(e.ErrorResponse[any](500, "Gagal mengambil data mapel", nil))
		}

		res := MapelResponse{
			ID:   created.ID.String(),
			Name: created.Name,
			Jurusan: jurusan.JurusanResponse{
				JurusanID: created.Jurusan.ID.String(),
				Name:      created.Jurusan.Name,
			},
		}

		return c.JSON(e.SuccessResponse(&res))
	}
}
