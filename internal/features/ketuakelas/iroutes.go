package ketuakelas

import (
	"monitoring-guru/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var KetuaKelasServ *KetuaKelasService

func RegisterRoutes(api fiber.Router, db *gorm.DB) {
	
	KetuaKelasServ = &KetuaKelasService{DB: db}
	handler := KetuaKelasHandler{Service: KetuaKelasServ}
	
	ketuaGroup := api.Group("/ketua-kelas")
	ketuaGroup.Post("/", middleware.JWTRoleProtected("kepala_sekolah"), handler.RegisterKetua())	
	ketuaGroup.Get("/", handler.GetAllKetuaKelasHandler())
	ketuaGroup.Get("/profile", middleware.JWTProtected(), handler.GetProfileHandler())
	ketuaGroup.Get("/unsigned", middleware.JWTProtected(), handler.GetUnsignedKetuaKelasHandler())
	ketuaGroup.Put("/", middleware.JWTRoleProtected("kepala_sekolah"), handler.UpdateKetuaKelasHandler())
	ketuaGroup.Delete("/:id", middleware.JWTRoleProtected("kepala_sekolah"), handler.DeleteKetuaHandler())
}
