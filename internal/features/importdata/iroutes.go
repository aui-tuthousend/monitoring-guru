package importdata

import (
	// "monitoring-guru/middleware"

	"github.com/gofiber/fiber/v2"
	// "gorm.io/gorm"
)

func RegisterRoutes(api fiber.Router) {
	
	api.Post("/import/jurusan", UploadJurusanHandler)
	api.Post("/import/guru", UploadGuruHandler)
}
