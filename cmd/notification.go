package main

import "sync"

type NotificationManager struct{
	clients map[string]map[chan string]bool
	mu sync.RWMutex   // read write mutex (lock)
}


// example data 
// clients := map[string]map[chan string]bool{
// 	"client1": {
// 		"channel1": true, OR 0x00000a40000:true
// 		"channel2": true, OR 0x00000a40001:true
// 	},
// 	"client2": {
// 		"channel1": true,
// 		"channel2": true,
// 	},
// }

func NewNotificationManager() *NotificationManager{
	return &NotificationManager{
		clients: make(map[string]map[chan string]bool),

	}
}

func (n *NotificationManager) AddClient(key string, client chan string) {
	n.mu.Lock()
	defer n.mu.Unlock()

	if n.clients[key] == nil {
		n.clients[key] = make(map[chan string]bool)
	}
	n.clients[key][client] = true

}

func (n *NotificationManager) RemoveClient(key string, client chan string) {
	n.mu.Lock()
	defer n.mu.Unlock()

	if clients := n.clients[key]; clients != nil {
		delete(clients, client)
		if len(clients) == 0{
			delete(n.clients, key)
		}
		
	}

	close(client)

}

func (n *NotificationManager) Notify(key string, message string) {
	// n.mu.RLock()
	// defer n.mu.RUnlock()

	// for client := range n.clients[key] {
	// 	select {
	// 		case client <- message:
	// 		default:
	// 			n.RemoveClient(key, client)
	// 	}
	// }

	// if clients := n.clients[key]; clients != nil {
	// 	for client := range clients {
	// 		client <- message
	// 	}
	// }
	
	n.mu.RLock()
	defer n.mu.RUnlock()

	for client := range n.clients[key] {
		select {
		case client <- message:
		default:
			n.RemoveClient(key, client)
		}
	}
}