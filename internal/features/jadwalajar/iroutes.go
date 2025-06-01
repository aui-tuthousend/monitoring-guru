package jadwalajar

import (
	"monitoring-guru/infrastructure/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(api fiber.Router, db *gorm.DB) {

	service := JadwalajarService{DB: db}
	handler := JadwalajarHandler{Service: &service}

	group := api.Group("jadwalajar", middleware.JWTProtected())
	group.Get("/", handler.GetAllJadwalAjar())
	group.Post("/", handler.CreateJadwalAjar())
	group.Get("/:id", handler.GetJadwalajarByID())
	group.Put("/", handler.UpdateJadwalajar())
	group.Get("/guru", handler.GetJadwalAjarByIDGuru())
	group.Get("/kelas", handler.GetJadwalAjarByIDKelas())
}
