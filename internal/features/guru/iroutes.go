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

	guruGroup := api.Group("/guru", middleware.JWTProtected())
	guruGroup.Get("/", middleware.JWTRoleProtected("kepala_sekolah"), handler.GetAllGuruHandler())
	guruGroup.Put("/", handler.UpdateGuruHandler())
	guruGroup.Post("/", handler.RegisterGuru())
	guruGroup.Get("/profile", handler.GetProfileHandler())
	guruGroup.Delete("/:id", handler.DeleteGuruHandler())
}
