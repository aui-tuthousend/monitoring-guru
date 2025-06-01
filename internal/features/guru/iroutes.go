package guru

import (
	"monitoring-guru/infrastructure/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(api fiber.Router, db *gorm.DB) {

	service := GuruService{DB: db}
	handler := GuruHandler{Service: &service}

	guruGroup := api.Group("/guru")
	guruGroup.Get("/", middleware.JWTProtected(), middleware.JWTRoleProtected("kepala_sekolah"), handler.GetAllGuruHandler())
	guruGroup.Put("/", middleware.JWTProtected(), handler.UpdateGuruHandler())
	guruGroup.Post("/", handler.RegisterGuru())
	guruGroup.Get("/profile", middleware.JWTProtected(), handler.GetProfileHandler())
	guruGroup.Delete("/:id", middleware.JWTProtected(), middleware.JWTRoleProtected("kepala_sekolah"), handler.DeleteGuruHandler())
}
