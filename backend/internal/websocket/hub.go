package websocket

import "sync"

// NewHub creates and returns a new Hub instance with initialized channels and maps.
type Hub struct {
	mu            sync.RWMutex
	clients       map[*Client]bool
	clientMap     map[string]*Client
	tickerMap     map[string][]*Client // ticker : [clients]
	activeMonitor map[string]bool      // ticker : is actively being monitored
	register      chan *Client
	unregister    chan *Client

}

func NewHub() *Hub {
	return &Hub{
		clients:       make(map[*Client]bool), // is use to register and unregister client 
		clientMap:     make(map[string]*Client), 
		tickerMap:     make(map[string][]*Client),
		activeMonitor: make(map[string]bool), // Initialize the active monitoring map
		register:      make(chan *Client),
		unregister:    make(chan *Client),
	}
}

func (h *Hub) Run() { 
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client] = true
			h.mu.Unlock()

		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)

				// Clean up corresponding mappings in clientMap and tickerMap
				for alertID, c := range h.clientMap {
					if c == client {
						delete(h.clientMap, alertID)
						// Clean up corresponding ticker
						for ticker, clients := range h.tickerMap {
							// Check if this client is part of the ticker's client slice
							for i, c := range clients {
								if c == client {
									// Remove the client from the tickerMap slice
									h.tickerMap[ticker] = append(clients[:i], clients[i+1:]...)
									break
								}
							}
						}
					}
				}
				close(client.send)
			}
			h.mu.Unlock()
		}
	}
}
