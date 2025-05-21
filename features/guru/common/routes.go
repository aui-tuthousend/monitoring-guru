package common

import (
	"monitoring-guru/features/guru/create"
	"monitoring-guru/features/guru/delete"
	getall "monitoring-guru/features/guru/getAll"
	getprofile "monitoring-guru/features/guru/getProfile"
	"monitoring-guru/features/guru/update"
	"monitoring-guru/infrastructure/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(api fiber.Router, db *gorm.DB) {
	guruGroup := api.Group("/guru")
	guruGroup.Post("/register", create.RegisterGuru(db))

	guruGroup.Get("/", middleware.JWTProtected(), middleware.JWTRoleProtected("kepala_sekolah"), getall.GetAllGuruHandler(db))
	guruGroup.Delete("/:id", middleware.JWTProtected(), middleware.JWTRoleProtected("kepala_sekolah"), delete.DeleteGuruHandler(db))
	guruGroup.Get("/profile", middleware.JWTProtected(), getprofile.GetProfileHandler(db))
	guruGroup.Put("/:id", middleware.JWTProtected(), update.UpdateGuruHandler(db))
}
