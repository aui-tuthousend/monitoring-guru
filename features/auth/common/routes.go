package common

import (
    
    "monitoring-guru/features/auth/login"
    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
)

func RegisterRoutes(api fiber.Router, db *gorm.DB) {
    // if all routes should be authenticated
    // group := api.Group("/routes", middleware.JWTProtected())
    group := api.Group("/auth")
    group.Post("/login", login.Login(db))
}
