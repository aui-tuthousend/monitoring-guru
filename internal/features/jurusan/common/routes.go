package common

import (
	"monitoring-guru/internal/features/jurusan/create"
	"monitoring-guru/internal/features/jurusan/delete"
	"monitoring-guru/internal/features/jurusan/getall"
	"monitoring-guru/internal/features/jurusan/getbyid"
	"monitoring-guru/internal/features/jurusan/update"
	"monitoring-guru/infrastructure/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(api fiber.Router, db *gorm.DB) {
	group := api.Group("jurusan", middleware.JWTProtected())
	group.Post("/", create.CreateJurusan(db))
	group.Get("/", getall.GetAllJurusan(db))
	group.Get("/:id", getbyid.GetJurusanByID(db))
	group.Put("/:id", update.UpdateJurusan(db))
	group.Delete("/:id", delete.DeleteJurusan(db))
}
