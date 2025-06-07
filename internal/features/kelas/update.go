package kelas

import (
	"fmt"
	e "monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// UpdateKelasRequest
// @Description Update kelas request body
type UpdateKelasRequest struct {
	// @Description ID of the kelas
	// @Required true
	// @Example "123456789"
	ID string `json:"id"`
	// @Description Name of the kelas
	// @Required true
	// @Example "XII RPL 1"
	Name string `json:"name"`
	// @Description Jurusan ID of the kelas
	// @Required true
	// @Example "123456789"
	// JurusanID string `json:"jurusan_id"`
	// @Description Ketua ID of the kelas
	// @Required true
	// @Example "123456789"
	KetuaID string `json:"ketua_id"`
	// @Description Wakil ID of the kelas
	// @Required true
	// @Example "123456789"
	// WakilID string `json:"wakil_id"`
	// @Description Is active of the kelas
	// @Required true
	// @Example true
	// IsActive bool `json:"is_active"`
}

// UpdateKelasHandler godoc
// @Summary Update kelas data
// @Description Update a kelas by ID
// @Tags Kelas
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body UpdateKelasRequest true "Update kelas request body"
// @Success 200 {object} KelasResponseWrapper
// @Failure 400 {object} KelasResponseWrapper
// @Failure 500 {object} KelasResponseWrapper
// @Router /api/kelas [put]
func (h *KelasHandler) UpdateKelasHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req UpdateKelasRequest
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

		// Parse UUIDs
		kelasID, err := parseUUID(req.ID, "ID")
		if err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, err.Error(), nil))
		}

		// jurusanID, err := parseUUID(req.JurusanID, "JurusanID")
		// if err != nil {
		// 	return c.Status(400).JSON(e.ErrorResponse[any](400, err.Error(), nil))
		// }

		ketuaID, err := parseUUID(req.KetuaID, "KetuaID")
		if err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, err.Error(), nil))
		}

		// wakilID, err := parseUUID(req.WakilID, "WakilID")
		// if err != nil {
		// 	return c.Status(400).JSON(e.ErrorResponse[any](400, err.Error(), nil))
		// }

		ketua, err := h.KetuaKelasService.GetKetuaKelas(ketuaID.String())
		if err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, "Ketua Kelas not found", nil))
		}
		// Update fields
		kelas := e.Kelas{
			ID:        kelasID,
			KetuaID:   ketuaID,
			// WakilID:   wakilID,
			// JurusanID: jurusanID,
			Name:      req.Name,
			// IsActive:  req.IsActive,
		}

		if err := h.Service.UpdateKelas(&kelas); err != nil {
			return c.Status(500).JSON(e.ErrorResponse[any](500, err.Error(), nil))
		}

		return c.JSON(e.SuccessResponse(&KelasResponse{
			ID:   kelas.ID.String(),
			Name: kelas.Name,
			// Jurusan: h.JurusanService.ResponseJurusanMapper(jurusan),
			KetuaKelas: h.KetuaKelasService.ResponseKetuaKelasMapper(ketua),
		}))
	}
}
