package kelas

import (
	"monitoring-guru/infrastructure/middleware"
	"monitoring-guru/internal/features/jurusan"
	"monitoring-guru/internal/features/ketuakelas"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(api fiber.Router, db *gorm.DB) {
	service := KelasService{DB: db}
	ketuaKelasService := ketuakelas.KetuaKelasService{DB: db}
	jurusanService := jurusan.JurusanService{DB: db}
	handler := KelasHandler{Service: &service, KetuaKelasService: &ketuaKelasService, JurusanService: &jurusanService}

	group := api.Group("kelas", middleware.JWTProtected())
	group.Post("/", handler.CreateKelas())
	group.Put("/", handler.UpdateKelasHandler())
	group.Get("/", handler.GetAllKelas())
	group.Get("/:id", handler.GetKelasByIDHandler())
	group.Get("/jurusan/:id", handler.GetKelasByJurusanHandler())
	group.Delete("/:id", handler.DeleteKelasHandler())
}
