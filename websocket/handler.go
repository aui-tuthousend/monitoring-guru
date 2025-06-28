package websocket

import (
	"encoding/json"
	"log"
	"monitoring-guru/internal/features/absenmasuk"
	"monitoring-guru/internal/features/jadwalajar"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func SetupWebSocket(app *fiber.App) {

	var WebsocketServ *WebsocketService = &WebsocketService{
		JadwalajarService: jadwalajar.JadwalajarServ,
		AbsenMasukService: absenmasuk.AbsenMasukServ,
	}
	
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

			}
			
			if message.Type == "update-kelas" {
				WebsocketServ.CreateAbsenMasuk(message.Payload)
			}	
		}
	}))	
	
}
