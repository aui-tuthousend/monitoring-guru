package login

import (
	r "monitoring-guru/infrastructure/repositories/auth"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Login godoc
//	@Summary
//	@Description	Log In
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		AuthRequest	true	"Create user request body"
//	@Success		200		{object}	AuthResponse
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/api/auth/login [post]
func Login(db *gorm.DB) fiber.Handler {
    return func(c *fiber.Ctx) error {
        var req AuthRequest
        if err := c.BodyParser(&req); err != nil {
            return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
        }

        user, err := r.FindUserByEmail(db, req.Email)
        if err != nil {
            return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
        }

        if user == nil || !r.CheckPasswordHash(req.Password, user.Password) {
            return c.Status(400).JSON(fiber.Map{"error": "Invalid credentials"})
        }

        token, _ := r.GenerateJWT(user.ID)
        return c.JSON(AuthResponse{Token: token})
    }
}
