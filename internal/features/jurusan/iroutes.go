package jurusan

import (
	"monitoring-guru/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var JurusanServ *JurusanService

func RegisterRoutes(api fiber.Router, db *gorm.DB) {

	JurusanServ = &JurusanService{DB: db}
	handler := JurusanHandler{Service: JurusanServ}

	group := api.Group("jurusan", middleware.JWTProtected())
	group.Post("/", handler.CreateJurusan())
	group.Get("/", handler.GetAllJurusan())
	group.Get("/:id", handler.GetJurusanByID())
	group.Put("/", handler.UpdateJurusan())
	group.Delete("/:id", handler.DeleteJurusan())
}
