package ruangan

import (
	"monitoring-guru/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var RuanganServ *RuanganService

func RegisterRoutes(api fiber.Router, db *gorm.DB) {

	RuanganServ = &RuanganService{DB: db}
	handler := RuanganHandler{Service: RuanganServ}

	group := api.Group("ruangan", middleware.JWTProtected())
	group.Post("/", handler.CreateRuangan())
	group.Get("/", handler.GetAllRuangan())
	group.Get("/:id", handler.GetRuanganByID())
	group.Put("/", handler.UpdateRuangan())
	group.Delete("/:id", handler.DeleteRuangan())
}
