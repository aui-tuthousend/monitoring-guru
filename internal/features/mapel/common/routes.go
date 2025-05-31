package common

import (
	"monitoring-guru/internal/features/mapel/create"
	"monitoring-guru/internal/features/mapel/delete"
	"monitoring-guru/internal/features/mapel/getAll"
	"monitoring-guru/internal/features/mapel/getbyid"
	"monitoring-guru/internal/features/mapel/update"
	"monitoring-guru/infrastructure/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(api fiber.Router, db *gorm.DB) {
	group := api.Group("mapel", middleware.JWTProtected())
	group.Post("/", create.CreateMapel(db))
	group.Get("/", getall.GetAllMapel(db))
	group.Get("/:id", getbyid.GetMapelByID(db))
	group.Put("/:id", update.UpdateMapel(db))
	group.Delete("/:id", delete.DeleteMapel(db))
}
