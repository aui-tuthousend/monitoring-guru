package common

import (
	"monitoring-guru/internal/features/jadwalajar/create"
	"monitoring-guru/internal/features/jadwalajar/getall"
	"monitoring-guru/internal/features/jadwalajar/getbyid"
	"monitoring-guru/internal/features/jadwalajar/getbyidguru"
	"monitoring-guru/internal/features/jadwalajar/getbyidkelas"
	"monitoring-guru/internal/features/jadwalajar/update"
	"monitoring-guru/infrastructure/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(api fiber.Router, db *gorm.DB) {
	group := api.Group("jadwalajar", middleware.JWTProtected())
	group.Get("/", getall.GetAllJadwalAjar(db))
	group.Get("/:id", getbyid.GetJadwalajarByID(db))
	group.Post("/", create.CreateJadwalAjar(db))
	group.Put("/", update.UpdateJadwalAjar(db))
	group.Get("/guru", getbyidguru.GetJadwalAjarByIDGuru(db))
	group.Get("/kelas", getbyidkelas.GetJadwalAjarByIDKelas(db))
}
