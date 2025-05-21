package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	auth "monitoring-guru/features/auth/common"
	guru "monitoring-guru/features/guru/common"
	ketua "monitoring-guru/features/ketuaKelas/common"
	user "monitoring-guru/features/user/common"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api")

	auth.RegisterRoutes(api, db)
	user.RegisterRoutes(api, db)
	guru.RegisterRoutes(api, db)
	ketua.RegisterRoutes(api, db)
}
