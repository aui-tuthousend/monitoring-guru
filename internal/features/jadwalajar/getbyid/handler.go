package getbyid

import (
	e "monitoring-guru/entities"
	r "monitoring-guru/infrastructure/repositories/jadwalajar"


	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GetJadwalajarByID godoc
// @Summary Get jadwalajar by ID
// @Description Get a jadwalajar by its ID
// @Tags Jadwalajar
// @Accept json
// @Produce json
// @Param id path string true "Jadwalajar ID"
// @Success 200 {object} GetJadwalajarByIDResponseWrapper
// @Failure 400 {object} GetJadwalajarByIDResponseWrapper
// @Failure 500 {object} GetJadwalajarByIDResponseWrapper
// @Router /api/jadwalajar/{id} [get]
func GetJadwalajarByID(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		jadwalajarID, err := uuid.Parse(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(e.ErrorResponse[any](400, "Invalid GUID format", nil))
		}

		jadwalajar, err := r.GetJadwalajarByID(db, jadwalajarID.String())
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(e.ErrorResponse[any](500, "Jadwalajar not found", nil))
		}

		return c.JSON(e.SuccessResponse(&jadwalajar))
	}
}
