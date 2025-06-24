package statuskelas

import (
	"monitoring-guru/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var StatusKelasServ *StatusKelasService

func RegisterRoutes(api fiber.Router, db *gorm.DB) {
	StatusKelasServ = &StatusKelasService{DB: db}
	handler := StatusKelasHandler{Service: StatusKelasServ}

	group := api.Group("statuskelas", middleware.JWTProtected())
	group.Get("/", handler.GetAllStatusKelas())
}