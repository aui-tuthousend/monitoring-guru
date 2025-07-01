package websocket

import (
	"encoding/json"
	"log"
	"monitoring-guru/internal/features/absenkeluar"
	"monitoring-guru/internal/features/absenmasuk"
	"monitoring-guru/internal/features/izin"
	"monitoring-guru/internal/features/jadwalajar"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func SetupWebSocket(app *fiber.App) {

	var WebsocketServ *WebsocketService = &WebsocketService{
		JadwalajarService: jadwalajar.JadwalajarServ,
		AbsenMasukService: absenmasuk.AbsenMasukServ,
		AbsenKeluarService: absenkeluar.AbsenKeluarServ,
		IzinService: izin.IzinServ,
	}
	
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/:group/:id", websocket.New(func(c *websocket.Conn) {
		group := c.Params("group")
		id := c.Params("id") 
		
		AddClientToGroup(group, id, c)
		log.Printf("Client %s joined group %s", id, group)
	
		defer func() {
			RemoveClientFromGroup(group, id)
			c.Close()
			log.Printf("Client %s disconnected from group %s", id, group)
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
			
			if message.Type == "clock-in" {
				WebsocketServ.CreateAbsenMasuk(message.Payload)
			} else if message.Type == "clock-out" {
				WebsocketServ.CreateAbsenKeluar(message.Payload)
			} else if message.Type == "create-izin" {
				WebsocketServ.CreateIzin(message.Payload)
			}
		}
	}))	

	
}
