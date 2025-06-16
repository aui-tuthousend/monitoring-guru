package jurusan

import (
	"monitoring-guru/middleware"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var JurusanServ *JurusanService

func RegisterRoutes(api fiber.Router, db *gorm.DB) {

	JurusanServ = &JurusanService{DB: db}
	handler := JurusanHandler{Service: JurusanServ}

	group := api.Group("jurusan", middleware.JWTProtected())
	group.Post("/", middleware.JWTRoleProtected("kepala_sekolah"),handler.CreateJurusan())
	group.Get("/", handler.GetAllJurusan())
	group.Get("/:id", handler.GetJurusanByID())
	group.Put("/", middleware.JWTRoleProtected("kepala_sekolah"),handler.UpdateJurusan())
	group.Delete("/:id", middleware.JWTRoleProtected("kepala_sekolah"),handler.DeleteJurusan())
}
