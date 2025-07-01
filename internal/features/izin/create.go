package izin

import (
	"time"

	e "monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// CreateIzinRequest
// @Description Create izin request body
type CreateIzinRequest struct {
	// @Description Tittle izin
	// @Required true
	// @Example "Izin"
	Judul string `json:"judul"`

	// @Description UUID jadwal ajar yang akan di-izin-kan
	// @Required true
	// @Example "a1b2c3d4-5e6f-7a8b-9c0d-e1f2a3b4c5d6"
	JadwalAjarID string `json:"jadwal_ajar_id"`

	// @Description Alasan izin guru
	// @Required true
	// @Example "Sakit Hati"
	Pesan string `json:"pesan"`
}

// RegisterIzin godoc
// @Summary Create izin guru
// @Description Create Izin request body
// @Tags Izin
// @Accept json
// @Produce json
// @Param request body CreateIzinRequest true "Create izin request body"
// @Success 200 {object} IzinResponseWrapper
// @Failure 400 {object} IzinResponseWrapper
// @Failure 500 {object} IzinResponseWrapper
// @Router /api/izin [post]
func (h *IzinHandler) CreateIzin() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req CreateIzinRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(e.ErrorResponse[any](400, err.Error(), nil))
		}

		if req.JadwalAjarID == "" || len(req.Pesan) < 3 || len(req.Judul) < 3 {
			return c.Status(400).JSON(e.ErrorResponse[any](400, "Invalid input data", nil))
		}

		loc, _ := time.LoadLocation("Asia/Jakarta")
		izinEntity := e.Izin{
			ID:           uuid.New(),
			// GuruID:       uuid.MustParse(req.GuruID),
			Judul:       req.Judul,
			Pesan:        req.Pesan,
			JadwalAjarID: uuid.MustParse(req.JadwalAjarID),
			TanggalIzin:  time.Now().In(loc),
			JamIzin: time.Now().In(loc).Format("15:04"),
			Approval:     false,
			Read: false,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		if err := h.Service.CreateIzin(&izinEntity); err != nil {
			return c.Status(500).JSON(e.ErrorResponse[any](500, "Gagal menyimpan izin", nil))
		}

		resp := h.Service.ResponseIzinMapper(&izinEntity)
		return c.JSON(e.SuccessResponse(resp))
	}
}
