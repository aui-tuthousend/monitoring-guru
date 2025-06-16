package jadwalajar

import (
	"monitoring-guru/middleware"
	"monitoring-guru/internal/features/guru"
	"monitoring-guru/internal/features/kelas"
	"monitoring-guru/internal/features/mapel"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var JadwalajarServ *JadwalajarService

func RegisterRoutes(api fiber.Router, db *gorm.DB) {

	JadwalajarServ = &JadwalajarService{DB: db}
	handler := JadwalajarHandler{
		Service: JadwalajarServ,
		GuruService: guru.GuruServ,
		MapelService: mapel.MapelServ,
		KelasService: kelas.KelasServ,
	}

	group := api.Group("jadwalajar", middleware.JWTProtected())
	group.Get("/", handler.GetAllJadwalAjar())
	group.Post("/", handler.CreateJadwalAjar())
	group.Get("/:id", handler.GetJadwalajarByID())
	group.Put("/", handler.UpdateJadwalajar())
	group.Get("/guru/:id/:hari", handler.GetJadwalAjarByIDGuru())
	group.Get("/kelas/:id/:hari", handler.GetJadwalAjarByIDKelas())
}
