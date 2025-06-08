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
	ketuaGroup.Post("/", handler.RegisterKetua())	
	ketuaGroup.Get("/", handler.GetAllKetuaKelasHandler())
	ketuaGroup.Get("/profile", middleware.JWTProtected(), handler.GetProfileHandler())
	ketuaGroup.Put("/", middleware.JWTProtected(), handler.UpdateKetuaKelasHandler())
	ketuaGroup.Delete("/:id", middleware.JWTProtected(), handler.DeleteKetuaHandler())
}
