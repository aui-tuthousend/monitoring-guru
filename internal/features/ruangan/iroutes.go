package ruangan

import (
	"monitoring-guru/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(api fiber.Router, db *gorm.DB) {
	service := RuanganService{DB: db}
	handler := RuanganHandler{Service: &service}

	group := api.Group("ruangan", middleware.JWTProtected())
	group.Post("/", handler.CreateRuangan())
	group.Get("/", handler.GetAllRuangan())
	group.Get("/:id", handler.GetRuanganByID())
	group.Put("/", handler.UpdateRuangan())
	group.Delete("/:id", handler.DeleteRuangan())
}
