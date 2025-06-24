package main

import (
	"monitoring-guru/internal/database"
	"monitoring-guru/routes"
	"monitoring-guru/websocket"

	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	app := fiber.New(fiber.Config{
		// EnablePrintRoutes: true,
	})
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		// AllowOrigins: "http://localhost:3000",
		AllowOrigins: "*",
		AllowHeaders: "Content-Type, Authorization",         
	}))
	
	db := database.Connect()
	routes.SetupRoutes(app, db)
	// websocket.SetupWebSocket(app, db)
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
