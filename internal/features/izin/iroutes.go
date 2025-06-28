package izin

import (
	"monitoring-guru/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var IzinServ *IzinService

func RegisterRoutes(api fiber.Router, db *gorm.DB) {

	IzinServ = &IzinService{DB: db}
	handler := IzinHandler{Service: IzinServ}

	izinGroup := api.Group("/izin", middleware.JWTProtected())
	izinGroup.Get("/", middleware.JWTRoleProtected("kepala_sekolah"), handler.GetAllIzinHandler())
	izinGroup.Get("/:id", middleware.JWTRoleProtected("kepala_sekolah"), handler.GetIzinByID())
	izinGroup.Put("/", handler.UpdateIzinHandler())
	izinGroup.Post("/", handler.CreateIzin())
	izinGroup.Delete("/:id", handler.DeleteIzinHandler())
}
