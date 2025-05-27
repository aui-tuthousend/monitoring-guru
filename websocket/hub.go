package websocket

import (
	"log"
	"sync"

	"github.com/gofiber/contrib/websocket"
)

var wsClients = make(map[string]*websocket.Conn)
var wsMutex sync.Mutex

func AddClient(id string, conn *websocket.Conn) {
	wsMutex.Lock()
	wsClients[id] = conn
	wsMutex.Unlock()
}

func RemoveClient(id string) {
	wsMutex.Lock()
	delete(wsClients, id)
	wsMutex.Unlock()
	log.Printf("client removed: %s logged out", id)
}

func SendToClient(id string, message string) {
	wsMutex.Lock()
	defer wsMutex.Unlock()

	if conn, ok := wsClients[id]; ok {
		if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			log.Printf("Failed to send message to %s: %v", id, err)
		}
	}
}

func CleanupWebSocketClients() {
	wsMutex.Lock()
	defer wsMutex.Unlock()

	for id, conn := range wsClients {
		log.Printf("Closing connection for client: %s", id)
		err := conn.Close()
		if err != nil {
			log.Printf("Error closing connection for client %s: %v", id, err)
		}
		delete(wsClients, id)
	}
}


func BroadcastToAll(message string, ch chan<-string, wg *sync.WaitGroup) {
	wsMutex.Lock()
	defer wsMutex.Unlock()

	for id, conn := range wsClients {
		log.Printf("broadcasting to %s with message: %s", id, message)
		
		wg.Add(1)
		go func(id string, conn *websocket.Conn) {
			defer wg.Done()
			if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
				log.Printf("error sending to %s: %v", id, err)
			}
		}(id, conn)
	}
	wg.Wait()
}

