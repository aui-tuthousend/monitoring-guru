package auth

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var AuthServ *AuthService

func RegisterRoutes(api fiber.Router, db *gorm.DB) {
	AuthServ = &AuthService{DB: db}
	handler := AuthHandler{Service: AuthServ}
	// if all routes should be authenticated
	// group := api.Group("/routes", middleware.JWTProtected())
	group := api.Group("/auth")
	// group.Post("/login", login.Login(db))
	group.Post("/login-guru", handler.LoginGuru())
	group.Post("/login-ketua-kelas", handler.LoginKetuaKelas())
}
