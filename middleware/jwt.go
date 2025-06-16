package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func JWTProtected() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(os.Getenv("JWT_SECRET")),
		ContextKey:   "user", // penting! ini yang kamu ambil di c.Locals("user")
		TokenLookup:  "header:Authorization",
		AuthScheme:   "Bearer",
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized",
	})
}

func JWTRoleProtected(requiredRole string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		role := claims["role"]
		userID := claims["sub"].(string) // atau apa pun key-nya
		c.Locals("userID", userID)

		if role != requiredRole {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Forbidden"})
		}
		return c.Next()
	}
}
