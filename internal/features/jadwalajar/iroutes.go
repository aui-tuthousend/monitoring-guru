package jadwalajar

import (
	"monitoring-guru/infrastructure/middleware"
	"monitoring-guru/internal/features/guru"
	"monitoring-guru/internal/features/kelas"
	"monitoring-guru/internal/features/mapel"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(api fiber.Router, db *gorm.DB) {

	service := JadwalajarService{DB: db}
	guruService := guru.GuruService{DB: db}
	mapelService := mapel.MapelService{DB: db}
	kelasService := kelas.KelasService{DB: db}
	handler := JadwalajarHandler{Service: &service, GuruService: &guruService, MapelService: &mapelService, KelasService: &kelasService}

	group := api.Group("jadwalajar", middleware.JWTProtected())
	group.Get("/", handler.GetAllJadwalAjar())
	group.Post("/", handler.CreateJadwalAjar())
	group.Get("/:id", handler.GetJadwalajarByID())
	group.Put("/", handler.UpdateJadwalajar())
	group.Get("/guru", handler.GetJadwalAjarByIDGuru())
	group.Get("/kelas", handler.GetJadwalAjarByIDKelas())
}
