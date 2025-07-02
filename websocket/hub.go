package websocket

import (
	"log"
	"sync"

	"github.com/gofiber/contrib/websocket"
)

var (
	wsGroups = make(map[string]map[string]*websocket.Conn)
	wsMutex  sync.Mutex
	wg       sync.WaitGroup
)

func AddClientToGroup(group string, id string, conn *websocket.Conn) {
	wsMutex.Lock()
	defer wsMutex.Unlock()

	if wsGroups[group] == nil {
		wsGroups[group] = make(map[string]*websocket.Conn)
	}
	wsGroups[group][id] = conn
}

func RemoveClientFromGroup(group string, id string) {
	wsMutex.Lock()
	defer wsMutex.Unlock()

	if groupClients, ok := wsGroups[group]; ok {
		delete(groupClients, id)
		if len(groupClients) == 0 {
			delete(wsGroups, group) // Bersihkan grup jika kosong
		}
	}
}

func CleanupWebSocketClients() {
	wsMutex.Lock()
	defer wsMutex.Unlock()

	for group, clients := range wsGroups {
		for id, conn := range clients {
			log.Printf("Closing connection for client: %s in group: %s", id, group)
			if err := conn.Close(); err != nil {
				log.Printf("Error closing connection for client %s: %v", id, err)
			}
			delete(clients, id)
		}
		delete(wsGroups, group)
	}
}

func SendToUserInGroup(group string, userID string, message string) {
	wsMutex.Lock()
	clients, ok := wsGroups[group]
	if !ok {
		wsMutex.Unlock()
		log.Printf("Group '%s' not found", group)
		return
	}

	conn, exists := clients[userID]
	wsMutex.Unlock()

	if !exists {
		log.Printf("User '%s' not found in group '%s'", userID, group)
		return
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			log.Printf("Failed to send message to user '%s' in group '%s': %v", userID, group, err)
		}
	}()
	wg.Wait()
}

func BroadcastToGroup(group string, message string) {
	wsMutex.Lock()
	clients, ok := wsGroups[group]
	if !ok {
		wsMutex.Unlock()
		return
	}
	copyClients := make(map[string]*websocket.Conn, len(clients))
	for id, conn := range clients {
		copyClients[id] = conn
	}
	wsMutex.Unlock()

	for id, conn := range copyClients {
		wg.Add(1)
		go func(id string, conn *websocket.Conn) {
			defer wg.Done()
			if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
				log.Printf("failed to send to %s in group %s: %v", id, group, err)
			}
		}(id, conn)
	}
	wg.Wait()
}

