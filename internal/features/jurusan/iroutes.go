package jurusan

import (
	"monitoring-guru/infrastructure/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(api fiber.Router, db *gorm.DB) {

	service := JurusanService{DB: db}
	handler := JurusanHandler{Service: &service}

	group := api.Group("jurusan", middleware.JWTProtected())
	group.Post("/", handler.CreateJurusan())
	group.Get("/", handler.GetAllJurusan())
	group.Get("/:id", handler.GetJurusanByID())
	group.Put("/", handler.UpdateJurusan())
	group.Delete("/:id", handler.DeleteJurusan())
}
