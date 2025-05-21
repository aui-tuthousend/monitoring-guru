package create

import (
	"strings"
	"time"

	e "monitoring-guru/entities"
	r "monitoring-guru/infrastructure/repositories/guru"
	"monitoring-guru/infrastructure/repositories/user"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// RegisterGuru godoc
// @Summary Mendaftarkan guru baru
// @Description Create Guru request body
// @Tags guru
// @Accept json
// @Produce json
// @Param request body CreateGuruRequest true "Create guru request body"
// @Success 200 {object} CreateGuruResponseWrapper
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/guru/register [post]
func RegisterGuru(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req CreateGuruRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}

		if strings.TrimSpace(req.NIP) == "" || len(req.Password) < 6 || (req.Jabatan != "guru" && req.Jabatan != "kepala_sekolah") {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid input data"})
		}

		hashedPassword, _ := user.HashPassword(req.Password)

		guru := e.Guru{
			ID:        uuid.New(),
			Nip:       req.NIP,
			Name:      req.Nama,
			Password:  hashedPassword,
			Jabatan:   req.Jabatan,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := r.CreateGuru(db, &guru); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Guru dengan NIP tersebut sudah ada"})
		}

		return c.JSON(CreateGuruResponseWrapper{
			Code:    200,
			Message: "Berhasil mendaftarkan guru",
			Data:    CreateGuruResponse{NIP: guru.Nip, Nama: guru.Name, Jabatan: guru.Jabatan},
		})
	}
}
