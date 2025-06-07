package ketuakelas

import (
	e "monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// GetProfileHandler godoc
// @Summary Get Logged user profile
// @Tags Ketua Kelas
// @Security BearerAuth
// @Produce json
// @Success 200 {object} KetuaKelasResponseWrapper
// @Failure 404 {object} KetuaKelasResponseWrapper
// @Router /api/ketua-kelas/profile [get]
func (h *KetuaKelasHandler) GetProfileHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userToken := c.Locals("user").(*jwt.Token)
		claims := userToken.Claims.(jwt.MapClaims)
		userID := claims["sub"].(string)

		id, err := uuid.Parse(userID)
		if err != nil {
			return c.Status(404).JSON(e.ErrorResponse[any](404, "Invalid Guid", nil))
		}

		user, err := h.Service.GetKetuaKelas(id.String())
		if err != nil {
			return c.Status(404).JSON(e.ErrorResponse[any](404, "Ketua kelas tidak ditemukan", nil))
		}

		return c.JSON(e.SuccessResponse(h.Service.ResponseKetuaKelasMapper(user)))
	}
}