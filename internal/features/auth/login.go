package auth

import (
	"monitoring-guru/internal/features/guru"
	"monitoring-guru/internal/features/ketuakelas"

	"github.com/gofiber/fiber/v2"
)

// LoginRequestBody
// @Description Login user request body
type AuthRequest struct {
	// Your Email
	Email string `json:"email"`
	// Your Password
	Password string `json:"password"`
}

// LoginGuruRequestBody
// @Description Login guru request body

type AuthGuruRequest struct {
	// Your NIP
	NIP string `json:"nip"`
	// Your Password
	Password string `json:"password"`
}

// LoginKetuaKelasRequestBody
// @Description Login ketua kelas request body
type AuthKetuaKelasRequest struct {
	// Your NISN
	NISN string `json:"nisn"`
	// Your Password
	Password string `json:"password"`
}

// LoginGuru godoc
//
//	@Summary
//	@Description	Log In
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		AuthGuruRequest	true	"Login Guru request body"
//	@Success		200		{object}	AuthResponse
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/api/auth/login-guru [post]
func (h *AuthHandler) LoginGuru() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req AuthGuruRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}
		user, err := h.Service.FindGuruByNip(req.NIP)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
		}
		if user == nil || !h.Service.CheckPasswordHash(req.Password, user.Password) {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid credentials"})
		}
		token, _ := h.Service.GenerateJWT(user.ID, user.Jabatan)
		return c.JSON(AuthGuruResponse{Token: token, UserData: &guru.GuruResponse{
			ID:       user.ID.String(),
			Nip:      user.Nip,
			Name:     user.Name,
			Jabatan:  user.Jabatan,
		}})
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
func (h *AuthHandler) LoginKetuaKelas() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req AuthKetuaKelasRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}
		user, err := h.Service.FindKetuaKelasByNisn(req.NISN)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Internal server error"})
		}
		if user == nil || !h.Service.CheckPasswordHash(req.Password, user.Password) {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid credentials"})
		}
		token, _ := h.Service.GenerateJWT(user.ID, "ketua_kelas")
		return c.JSON(AuthKetuaKelasResponse{Token: token, UserData: &ketuakelas.KetuaKelasResponse{
			ID:       user.ID.String(),
			Nisn:     user.Nisn,
			Name:     user.Name,
		}})
	}
}
