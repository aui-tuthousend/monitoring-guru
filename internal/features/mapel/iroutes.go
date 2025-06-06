package mapel

import (
	"monitoring-guru/infrastructure/middleware"
	"monitoring-guru/internal/features/jurusan"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(api fiber.Router, db *gorm.DB) {
	service := MapelService{DB: db}
	jurusanService := jurusan.JurusanService{DB: db}
	handler := MapelHandler{Service: &service, JurusanService: &jurusanService}

	group := api.Group("mapel", middleware.JWTProtected())
	group.Post("/", handler.CreateMapel())
	group.Get("/", handler.GetAllMapel())
	group.Get("/:id", handler.GetMapelByID())
	group.Put("/", handler.UpdateMapel())
	group.Delete("/:id", handler.DeleteMapel())
}
