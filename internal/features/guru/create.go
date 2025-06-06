package guru

import (
	"strings"
	"time"

	e "monitoring-guru/entities"
	"monitoring-guru/infrastructure/repositories/user"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)


// CreateGuruRequestBody
// @Description Create guru request body
type CreateGuruRequest struct {
	// @Description NIP of the guru
	// @Required true
	// @Example "123456789"
	NIP string `json:"nip"`
	// @Description Name of the guru
	// @Required true
	// @Example "John Doe"
	Nama string `json:"nama"`
	// @Description Password of the guru
	// @Required true
	// @Example "password123"
	// @MinLength 6
	Password string `json:"password"`
	// @Description Jabatan of the guru
	// @Required true
	// @Enum "guru" "kepala_sekolah"
	// @Example "guru"
	Jabatan string `json:"jabatan"`
}


// RegisterGuru godoc
// @Summary Mendaftarkan guru baru
// @Description Create Guru request body
// @Tags Guru
// @Accept json
// @Produce json
// @Param request body CreateGuruRequest true "Create guru request body"
// @Success 200 {object} GuruResponseWrapper
// @Failure 400 {object} GuruResponseWrapper
// @Failure 500 {object} GuruResponseWrapper
// @Router /api/guru [post]
func (h *GuruHandler) RegisterGuru() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req CreateGuruRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, err.Error(), nil))
		}

		if strings.TrimSpace(req.NIP) == "" || len(req.Password) < 6 || (req.Jabatan != "guru" && req.Jabatan != "kepala_sekolah") {
			return c.Status(400).JSON(e.ErrorResponse[any](400, "Invalid input data", nil))
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

		if err := h.Service.CreateGuru(&guru); err != nil {
			return c.Status(500).JSON(e.ErrorResponse[any](500, "Guru dengan NIP tersebut sudah ada", nil))
		}

		return c.JSON(e.SuccessResponse(h.Service.ResponseGuruMapper(&guru)))
	}
}
