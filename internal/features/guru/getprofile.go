package guru


import (
	e "monitoring-guru/entities"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// Get Profile godoc
//	@Summary
//	@Description	Get Logged user profile
//	@Tags			Guru
//	@Produce		json
// @Security     BearerAuth
//	@Success		200		{object}	GuruResponseWrapper
//	@Failure		404		{object}	GuruResponseWrapper
//	@Router			/api/guru/profile [get]
func (h *GuruHandler) GetProfileHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userToken := c.Locals("user").(*jwt.Token)
		claims := userToken.Claims.(jwt.MapClaims)
		userID := claims["sub"].(string)

		id, err := uuid.Parse(userID)
		if err != nil {
			return c.Status(404).JSON(e.ErrorResponse[any](404, "Invalid Guid", nil))
		}

		user, err := h.Service.GetGuru(id.String())
		if err != nil {
			return c.Status(404).JSON(e.ErrorResponse[any](404, "User not found", nil))
		}

		return c.JSON(e.SuccessResponse(&GuruResponse{
			ID:      user.ID.String(),
			NIP:     user.Nip,
			Nama:    user.Name,
			Jabatan: user.Jabatan,
		}))
	}
}
