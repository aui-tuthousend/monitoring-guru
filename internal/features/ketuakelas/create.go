package ketuakelas

import (
	"strings"
	"time"

	e "monitoring-guru/entities"
	"monitoring-guru/infrastructure/repositories/user"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// CreateKetuaRequestBody
// @Description Create ketua request body
type CreateKetuaRequest struct {
	// @Description NISN of the ketua
	// @Required true
	// @Example "123456789"
	NISN string `json:"nisn"`
	// @Description Name of the ketua
	// @Required true
	// @Example "John Doe"
	Name string `json:"nama"`
	// @Description Password of the ketua
	// @Required true
	// @Example "password123"
	// @MinLength 6
	Password string `json:"password"`
}

// CreateKetuaRequest godoc
// @summary Mendaftarkan ketua kelas baru
// @Tags			Ketua Kelas
// @Accept			json
// @Produce			json
// @Param			request	body		CreateKetuaRequest	true	"Create ketua request body"
// @Success		200		{object}	KetuaKelasResponseWrapper
// @Failure		400		{object}	KetuaKelasResponseWrapper
// @Failure		500		{object}	KetuaKelasResponseWrapper
// @Router			/api/ketua-kelas [post]
func (h *KetuaKelasHandler) RegisterKetua() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req CreateKetuaRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, err.Error(), nil))
		}

		if strings.TrimSpace(req.NISN) == "" || len(req.Password) < 6 {
			return c.Status(400).JSON(e.ErrorResponse[any](400, "Invalid input", nil))
		}

		hashed, _ := user.HashPassword(req.Password)

		ketua := e.KetuaKelas{
			ID:        uuid.New(),
			Nisn:      req.NISN,
			Name:      req.Name,
			Password:  hashed,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := h.Service.CreateKetuaKelas(&ketua); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Ketua dengan NISN tersebut sudah ada"})
		}

		return c.JSON(e.SuccessResponse(h.Service.ResponseKetuaKelasMapper(&ketua)))
	}
}
