package jurusan

import (
	e "monitoring-guru/entities"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CreateJurusanRequest struct {
	Name string `json:"name" validate:"required"`
}

// CreateJurusan godoc
// @Summary Create jurusan
// @Description Create a new jurusan
// @Tags jurusan
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body CreateJurusanRequest true "Jurusan body"
// @Success 201 {object} JurusanResponseWrapper
// @Failure 400 {object} JurusanResponseWrapper
// @Failure 500 {object} JurusanResponseWrapper
// @Router /api/jurusan [post]
func (h *JurusanHandler) CreateJurusan() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req CreateJurusanRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(e.ErrorResponse[any](400, "Invalid request body", nil))
		}

		if req.Name == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Name is required",
			})
		}

		jurusan := e.Jurusan{
			ID:        uuid.New(),
			Name:      req.Name,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := h.Service.CreateJurusan(&jurusan); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(e.ErrorResponse[any](500, err.Error(), nil))
		}

		return c.JSON(e.SuccessResponse(h.Service.ResponseJurusanMapper(&jurusan)))
	}
}