package common

import (
	"monitoring-guru/features/ketuaKelas/create"
	"monitoring-guru/features/ketuaKelas/delete"
	getall "monitoring-guru/features/ketuaKelas/getAll"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(api fiber.Router, db *gorm.DB) {
	ketuaGroup := api.Group("/ketua-kelas")
	ketuaGroup.Post("/register", create.RegisterKetua(db))
	ketuaGroup.Get("/", getall.GetAllKetuaKelasHandler(db))
	ketuaGroup.Delete("/:id", delete.DeleteKetuaHandler(db))
}
