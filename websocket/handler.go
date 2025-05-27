package websocket

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var wg sync.WaitGroup

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
		defer func() {
			RemoveClient(id)
			c.Close()
		}()
		
		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				break
			}
	
			var payload struct {
				IsActive bool `json:"IsActive"`
				UUID string `json:"UUID"`
			}

			error := json.Unmarshal(msg, &payload)
			if error != nil {
				log.Println("error unmarshalling:", error)
			} else {
				log.Printf("Parsed payload: %+v\n", payload.IsActive)
			}
	
			if payload.IsActive {
				log.Printf("activating")
				// db.Model(&Class{}).Where("id = ?", payload.uuid).Update("is_active", true)
				BroadcastToAll(`{"UUID":"` + payload.UUID + `", "IsActive":true}`, make(chan<- string), &wg)
			}
	
			if !payload.IsActive {
				// db.Model(&Class{}).Where("id = ?", payload.uuid).Update("is_active", false)
				BroadcastToAll(`{"UUID":"` + payload.UUID + `", "IsActive":false}`, make(chan<- string), &wg)
			}
		}
	}))
	
}
