package ketuakelas

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(api fiber.Router, db *gorm.DB) {
	
	service := KetuaKelasService{DB: db}
	handler := KetuaKelasHandler{Service: &service}

	ketuaGroup := api.Group("/ketua-kelas")
	ketuaGroup.Post("/register", handler.RegisterKetua())
	ketuaGroup.Get("/", handler.GetAllKetuaKelasHandler())
	ketuaGroup.Delete("/:id", handler.DeleteKetuaHandler())
}
