package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	auth "monitoring-guru/features/auth/common"
	guru "monitoring-guru/features/guru/common"
	jurusan "monitoring-guru/features/jurusan/common"
	ketua "monitoring-guru/features/ketuaKelas/common"
	mapel "monitoring-guru/features/mapel/common"
	ruangan "monitoring-guru/features/ruangan/common"
	user "monitoring-guru/features/user/common"
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
}
