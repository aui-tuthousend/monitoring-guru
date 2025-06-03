package kelas

import (
	"monitoring-guru/infrastructure/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(api fiber.Router, db *gorm.DB) {
	service := KelasService{DB: db}
	handler := KelasHandler{Service: &service}

	group := api.Group("kelas", middleware.JWTProtected())
	group.Post("/", handler.CreateKelas())
}
