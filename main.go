package main

import (
	"monitoring-guru/internal/database"
	"monitoring-guru/routes"
	"monitoring-guru/websocket"
	e "monitoring-guru/entities"

	"log"
	"os"
	"os/signal"
	"syscall"

	"monitoring-guru/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

//	@title			Monitoring Guru
//	@version		1.0
//	@description	endpoints for monitoring guru XD
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
	db.AutoMigrate(&e.Guru{}, &e.Jurusan{}, &e.KetuaKelas{}, &e.Ruangan{}) //migrate later
	routes.SetupRoutes(app, db)
	websocket.SetupWebSocket(app, db)
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("What are you doing here?")
	})

	c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)

    go func() {
        <-c
        log.Println("Gracefully shutting down...")
        websocket.CleanupWebSocketClients()
        if err := app.Shutdown(); err != nil {
            log.Fatalf("Error shutting down server: %v", err)
        }
    }()

	app.Listen(":8080")
}
