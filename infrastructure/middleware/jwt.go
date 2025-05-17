package middleware

import (
    jwtware "github.com/gofiber/jwt/v3"
    "github.com/gofiber/fiber/v2"
    "os"
)

func JWTProtected() func(*fiber.Ctx) error {
    return jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
    return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
        "error": "Unauthorized",
    })
}
