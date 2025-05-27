package common

import (
	"monitoring-guru/features/jurusan/create"
	"monitoring-guru/features/jurusan/delete"
	getall "monitoring-guru/features/jurusan/getAll"
	"monitoring-guru/features/jurusan/getbyid"
	"monitoring-guru/features/jurusan/update"
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
