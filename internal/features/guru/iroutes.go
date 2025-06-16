package guru

import (
	"monitoring-guru/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var GuruServ *GuruService

func RegisterRoutes(api fiber.Router, db *gorm.DB) {

	GuruServ = &GuruService{DB: db}
	handler := GuruHandler{Service: GuruServ}

	guruGroup := api.Group("/guru")
	guruGroup.Get("/", middleware.JWTRoleProtected("kepala_sekolah"), handler.GetAllGuruHandler())
	guruGroup.Put("/", middleware.JWTRoleProtected("guru"), handler.UpdateGuruHandler())
	guruGroup.Post("/", middleware.JWTRoleProtected("kepala_sekolah"), handler.RegisterGuru())
	guruGroup.Get("/profile", middleware.JWTProtected(), handler.GetProfileHandler())
	guruGroup.Delete("/:id", middleware.JWTRoleProtected("kepala_sekolah"), handler.DeleteGuruHandler())
}
