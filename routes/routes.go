package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"monitoring-guru/internal/features/auth"
	"monitoring-guru/internal/features/guru"
	"monitoring-guru/internal/features/jadwalajar"
	"monitoring-guru/internal/features/jurusan"
	"monitoring-guru/internal/features/kelas"
	ketua "monitoring-guru/internal/features/ketuakelas"
	"monitoring-guru/internal/features/mapel"
	"monitoring-guru/internal/features/ruangan"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api")

	jurusan.RegisterRoutes(api, db)
	auth.RegisterRoutes(api, db)
	guru.RegisterRoutes(api, db)
	ketua.RegisterRoutes(api, db)
	kelas.RegisterRoutes(api, db)
	mapel.RegisterRoutes(api, db)
	jadwalajar.RegisterRoutes(api, db)
	ruangan.RegisterRoutes(api, db)
}
