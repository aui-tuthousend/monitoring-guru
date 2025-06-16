package kelas

import (
	"monitoring-guru/middleware"
	"monitoring-guru/internal/features/jurusan"
	"monitoring-guru/internal/features/ketuakelas"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var KelasServ *KelasService

func RegisterRoutes(api fiber.Router, db *gorm.DB) {

	KelasServ = &KelasService{DB: db}
	handler := KelasHandler{Service: KelasServ, KetuaKelasService: ketuakelas.KetuaKelasServ, JurusanService: jurusan.JurusanServ}

	group := api.Group("kelas")
	group.Post("/", middleware.JWTRoleProtected("kepala_sekolah"), handler.CreateKelas())
	group.Put("/", middleware.JWTRoleProtected("kepala_sekolah"), handler.UpdateKelasHandler())
	group.Get("/", handler.GetAllKelas())
	group.Get("/:id", handler.GetKelasByIDHandler())
	group.Get("/jurusan/:id", handler.GetKelasByJurusanHandler())
	group.Delete("/:id", middleware.JWTRoleProtected("kepala_sekolah"), handler.DeleteKelasHandler())
	group.Get("/ketua/:id", func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusServiceUnavailable, "On vacation!")
	})
}
