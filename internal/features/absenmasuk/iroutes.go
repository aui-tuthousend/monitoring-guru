package absenmasuk

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var AbsenMasukServ *AbsenMasukService

func RegisterRoutes(api fiber.Router, db *gorm.DB) {
	AbsenMasukServ = &AbsenMasukService{DB: db}
}