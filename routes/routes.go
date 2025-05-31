package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"monitoring-guru/internal/features/auth"
	"monitoring-guru/internal/features/guru"
	jurusan "monitoring-guru/internal/features/jurusan/common"
	ketua "monitoring-guru/internal/features/ketuaKelas/common"
	mapel "monitoring-guru/internal/features/mapel/common"
	ruangan "monitoring-guru/internal/features/ruangan/common"
	user "monitoring-guru/internal/features/user/common"
	jadwalajar "monitoring-guru/internal/features/jadwalajar/common"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api")

	auth.RegisterRoutes(api, db)
	user.RegisterRoutes(api, db)
	guru.RegisterRoutes(api, db)
	ketua.RegisterRoutes(api, db)
	ruangan.RegisterRoutes(api, db)
	mapel.RegisterRoutes(api, db)
	jurusan.RegisterRoutes(api, db)
	jadwalajar.RegisterRoutes(api, db)
}
