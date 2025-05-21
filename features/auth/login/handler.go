package login

import (
	r "monitoring-guru/infrastructure/repositories/auth"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// LoginGuru godoc
//
//	@Summary
//	@Description	Log In
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		AuthGuruRequest	true	"Create user request body"
//	@Success		200		{object}	AuthResponse
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/api/auth/login-guru [post]
func LoginGuru(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req AuthGuruRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}
		user, err := r.FindGuruByNip(db, req.NIP)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
		}
		if user == nil || !r.CheckPasswordHash(req.Password, user.Password) {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid credentials"})
		}
		token, _ := r.GenerateJWT(user.ID, user.Jabatan)
		return c.JSON(AuthResponse{Token: token})
	}
}

// LoginKetuaKelas godoc
//
//	@Summary
//	@Description	Log In
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		AuthKetuaKelasRequest	true	"Create user request body"
//	@Success		200		{object}	AuthResponse
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/api/auth/login-ketua-kelas [post]
func LoginKetuaKelas(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req AuthKetuaKelasRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}
		user, err := r.FindKetuaKelasByNisn(db, req.NISN)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
		}
		if user == nil || !r.CheckPasswordHash(req.Password, user.Password) {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid credentials"})
		}
		token, _ := r.GenerateJWT(user.ID, "ketua_kelas")
		return c.JSON(AuthResponse{Token: token})
	}
}
