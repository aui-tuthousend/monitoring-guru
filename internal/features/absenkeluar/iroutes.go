package absenkeluar

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var AbsenKeluarServ *AbsenKeluarService

func RegisterRoutes(api fiber.Router, db *gorm.DB) {
	AbsenKeluarServ = &AbsenKeluarService{DB: db}
}