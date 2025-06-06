package mapel

import (
	"fmt"
	e "monitoring-guru/entities"

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
// @Tags        Mapel
// @Security    BearerAuth
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

		jurusan, err := h.JurusanService.GetJurusanByID(jurusanID.String())
		if err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, "Jurusan not found", nil))
		}

		mapel := e.Mapel{
			ID:        uuid.New(),
			Name:      req.Name,
			JurusanID: jurusanID,
		}

		if err := h.Service.CreateMapel(&mapel); err != nil {
			return c.Status(500).JSON(e.ErrorResponse[any](500, err.Error(), nil))
		}

		res := MapelResponse{
			ID:   mapel.ID.String(),
			Name: mapel.Name,
			Jurusan: h.JurusanService.ResponseJurusanMapper(jurusan),
		}

		return c.JSON(e.SuccessResponse(&res))
	}
}
