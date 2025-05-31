package common

import (
	"monitoring-guru/internal/features/ketuaKelas/create"
	"monitoring-guru/internal/features/ketuaKelas/delete"
	getall "monitoring-guru/internal/features/ketuaKelas/getall"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(api fiber.Router, db *gorm.DB) {
	ketuaGroup := api.Group("/ketua-kelas")
	ketuaGroup.Post("/register", create.RegisterKetua(db))
	ketuaGroup.Get("/", getall.GetAllKetuaKelasHandler(db))
	ketuaGroup.Delete("/:id", delete.DeleteKetuaHandler(db))
}
