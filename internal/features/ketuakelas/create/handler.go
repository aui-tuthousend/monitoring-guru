package create

import (
	"strings"
	"time"

	e "monitoring-guru/entities"
	r "monitoring-guru/infrastructure/repositories/ketua"
	"monitoring-guru/infrastructure/repositories/user"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// CreateKetuaRequest godoc
// @summary Create Ketua Kelas request body
// @Description	Create Ketua Kelas baru request body
// @Tags			ketua kelas
// @Accept			json
// @Produce		json
// @Param			request	body		CreateKetuaRequest	true	"Create ketua request body"
// @Success		200		{object}	CreateKetuaResponseWrapper
// @Failure		400		{object}	map[string]string
// @Failure		500		{object}	map[string]string
// @Router			/api/ketua-kelas/register [post]
func RegisterKetua(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req CreateKetuaRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}

		if strings.TrimSpace(req.NISN) == "" || len(req.Password) < 6 {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
		}

		hashed, _ := user.HashPassword(req.Password)

		ketua := e.KetuaKelas{
			ID:        uuid.New(),
			Nisn:      req.NISN,
			Name:      req.Nama,
			Password:  hashed,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := r.CreateKetuaKelas(db, &ketua); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Ketua dengan NISN tersebut sudah ada"})
		}

		return c.JSON(CreateKetuaResponseWrapper{
			Code:    200,
			Message: "Berhasil mendaftarkan ketua kelas",
			Data:    CreateKetuaResponse{NISN: ketua.Nisn, Nama: ketua.Name},
		})
	}
}
