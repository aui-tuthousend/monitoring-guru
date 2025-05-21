package getprofile

import (
	"monitoring-guru/infrastructure/repositories/guru"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetProfileHandler(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userToken := c.Locals("user").(*jwt.Token)
		claims := userToken.Claims.(jwt.MapClaims)
		userID := claims["sub"].(string)

		id, err := uuid.Parse(userID)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Invalid Guid"})
		}

		user, err := guru.GetGuru(db, id.String())
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "User not found"})
		}

		return c.JSON(ProfileGuruResponse{
			Name:      user.Name,
			Nip:       user.Nip,
			Jabatan:   user.Jabatan,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
}
