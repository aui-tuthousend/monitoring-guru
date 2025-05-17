package common

import (
    "monitoring-guru/infrastructure/middleware"
    "monitoring-guru/features/user/create"
    "monitoring-guru/features/user/getProfile"
    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
)

func RegisterRoutes(api fiber.Router, db *gorm.DB) {
    // if all routes should be authenticated
    // group := api.Group("/routes", middleware.JWTProtected())

    group := api.Group("/user")
    group.Post("/register", create.Register(db))

    // if just one of them
    group.Get("/profile", middleware.JWTProtected(), getprofile.Profile(db))
}
