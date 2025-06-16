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
	
	ketuaGroup := api.Group("/ketua-kelas", middleware.JWTProtected())
	ketuaGroup.Post("/", handler.RegisterKetua())	
	ketuaGroup.Get("/", handler.GetAllKetuaKelasHandler())
	ketuaGroup.Get("/profile", handler.GetProfileHandler())
	ketuaGroup.Get("/unsigned", handler.GetUnsignedKetuaKelasHandler())
	ketuaGroup.Put("/", handler.UpdateKetuaKelasHandler())
	ketuaGroup.Delete("/:id", handler.DeleteKetuaHandler())
}
