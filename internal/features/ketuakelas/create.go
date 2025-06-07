package ketuakelas

import (
	"strings"
	"time"

	e "monitoring-guru/entities"
	"monitoring-guru/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// CreateKetuaRequestBody
// @Description Create ketua request body
// @Description NISN of the ketua
type CreateKetuaRequest struct {
	// @Description NISN of the ketua
	// @Required true
	// @Example "123456789"
	NISN string `json:"nisn"`
	// @Description Name of the ketua
	// @Required true
	// @Example "John Doe"
	Name string `json:"name"`
	// @Description Password of the ketua
	// @Required true
	// @Example "password123"
	// @MinLength 6
	Password string `json:"password"`
}

// CreateKetuaRequest godoc
// @summary Create Ketua Kelas request body
// @Description	Create Ketua Kelas baru request body
// @Tags			ketua kelas
// @Accept			json
// @Produce		json
// @Param			request	body		CreateKetuaRequest	true	"Create ketua request body"
// @Success		200		{object}	KetuaKelasResponseWrapper
// @Failure		400		{object}	map[string]string
// @Failure		500		{object}	map[string]string
// @Router			/api/ketua-kelas/register [post]
func (h *KetuaKelasHandler) RegisterKetua() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req CreateKetuaRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}

		if strings.TrimSpace(req.NISN) == "" || len(req.Password) < 6 {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
		}

		hashed, _ := utils.HashPassword(req.Password)

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

		return c.JSON(e.SuccessResponse(&KetuaKelasResponse{
			ID:   ketua.ID,
			Nisn: ketua.Nisn,
			Name: ketua.Name,
		}))
	}
}
