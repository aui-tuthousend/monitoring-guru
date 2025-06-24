package websocket

import (
	"encoding/json"
	"log"
	"monitoring-guru/entities"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)


func SetupWebSocket(app *fiber.App, db *gorm.DB) {
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
		id := c.Params("id")
		AddClient(id, c)
		log.Printf("Client connected: %s", id)
	
		defer func() {
			RemoveClient(id)
			c.Close()
			log.Printf("Client disconnected: %s", id)
		}()
	
		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				break
			}
	
			var message struct {
				Type    string          `json:"type"`
				Payload json.RawMessage `json:"payload"`
			}
			
			if err := json.Unmarshal(msg, &message); err != nil {
				log.Println("Error unmarshalling message:", err)
				continue
			}
			
			if message.Type == "update-kelas" {
				var payload struct {
					IsActive bool   `json:"is_active"`
					Id       string `json:"id"`
					Mapel    string `json:"mapel"`
					Pengajar string `json:"pengajar"`
					Ruangan  string `json:"ruangan"`
				}
			
				if err := json.Unmarshal(message.Payload, &payload); err != nil {
					log.Println("Error unmarshalling payload:", err)
					continue
				}
			
				if payload.Id == "" {
					log.Println("Payload missing ID")
					continue
				}
			
				log.Printf("Parsed payload: %+v\n", payload)
			
				err := db.Model(&entities.StatusKelas{}).Where("id = ?", payload.Id).Update("is_active", payload.IsActive).Error
				log.Printf("Attempting to update StatusKelas id=%s isActive=%v\n", payload.Id, payload.IsActive)

				if err != nil {
					log.Printf("Failed to update DB: %v", err)
				}
			
				response, _ := json.Marshal(struct {
					Type    string      `json:"type"`
					Payload interface{} `json:"payload"`
				}{
					Type:    "update-kelas",
					Payload: payload,
				})
			
				BroadcastToAll(string(response))
			}			
		}
	}))	
	
}
