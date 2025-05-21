package getprofile

import (
	r "monitoring-guru/infrastructure/repositories/user"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Get Profile godoc
//	@Summary
//	@Description	Get Logged user profile
//	@Tags			user
//	@Produce		json
// @Security     BearerAuth
//	@Success		200		{object}	GetUserResponse
//	@Failure		404		{object}	map[string]string
//	@Router			/api/user/profile [get]
func Profile(db *gorm.DB) fiber.Handler {
    return func(c *fiber.Ctx) error {
        userToken := c.Locals("user").(*jwt.Token)
        claims := userToken.Claims.(jwt.MapClaims)
        userID := claims["sub"].(string)

        id, err := uuid.Parse(userID)
        if err != nil {
            return c.Status(404).JSON(fiber.Map{"error": "Invalid Guid"})
        }

		user, err := r.GetUser(db, id.String())
        if err != nil {
            return c.Status(404).JSON(fiber.Map{"error": "User not found"})
        }

        return c.JSON(GetUserResponse{
            Name:  user.Name,
            Email: user.Email,
        })
    }
}

// 