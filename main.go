package main

import (
	"monitoring-guru/database"
	"monitoring-guru/routes"
	"monitoring-guru/docs"
	// e "monitoring-guru/entities"
	"os"

	_ "monitoring-guru/docs" // for Swagger docs

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

//	@title			your own API application
//	@version		1.0
//	@description	Restful API using go fiber
//	@host			127.0.0.1:8080
//	@BasePath		/

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
    if os.Getenv("ENV") != "production" {
        docs.SwaggerInfo.Host = "127.0.0.1:8080"
    } else {
        docs.SwaggerInfo.Host = "your-deployment" // change later
    }
    
    app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
	})
    app.Use(logger.New())
    database.Connect()
    db := database.DB
    // db.AutoMigrate(&e.User{}) //migrate later
    // db.AutoMigrate(&e.User{}, &e.otherEntities{})
    routes.SetupRoutes(app, db)
    app.Get("/swagger/*", swagger.HandlerDefault)

    app.Listen(":8080")
}
