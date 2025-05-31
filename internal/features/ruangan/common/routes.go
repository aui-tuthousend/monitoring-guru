package common

import (
	"monitoring-guru/internal/features/ruangan/create"
	"monitoring-guru/internal/features/ruangan/delete"
	getall "monitoring-guru/internal/features/ruangan/getAll"
	"monitoring-guru/internal/features/ruangan/getbyid"
	"monitoring-guru/internal/features/ruangan/update"
	"monitoring-guru/infrastructure/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(api fiber.Router, db *gorm.DB) {
	group := api.Group("ruangan", middleware.JWTProtected())
	group.Post("/", create.CreateRuangan(db))
	group.Get("/", getall.GetAllRuangan(db))
	group.Get("/:id", getbyid.GetRuanganByID(db))
	group.Put("/:id", update.UpdateRuangan(db))
	group.Delete("/:id", delete.DeleteRuangan(db))
}
