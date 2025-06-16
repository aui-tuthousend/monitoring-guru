package mapel

import (
	"monitoring-guru/middleware"
	"monitoring-guru/internal/features/jurusan"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var MapelServ *MapelService

func RegisterRoutes(api fiber.Router, db *gorm.DB) {
	MapelServ = &MapelService{DB: db}
	handler := MapelHandler{Service: MapelServ, JurusanService: jurusan.JurusanServ}

	group := api.Group("mapel", middleware.JWTRoleProtected("kepala_sekolah"))
	group.Post("/", handler.CreateMapel())
	group.Get("/", handler.GetAllMapel())
	group.Get("/:id", handler.GetMapelByID())
	group.Put("/", handler.UpdateMapel())
	group.Delete("/:id", handler.DeleteMapel())
}
