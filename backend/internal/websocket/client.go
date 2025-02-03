package websocket

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	// "github.com/mahirjain_10/stock-alert-app/backend/internal/types"
	"github.com/mahirjain_10/stock-alert-app/backend/internal/types"
	"github.com/mahirjain_10/stock-alert-app/backend/internal/utils"
)

const (
	writeWait      = 10 * time.Second    // Time to wait for writing a message
	pongWait       = 60 * time.Second    // Time to wait before considering the connection dead if no Pong is received
	pingPeriod     = (pongWait * 9) / 10 // Ping interval (90% of pongWait)
	maxMessageSize = 512                 // Max allowed message size for WebSocket
)

var (
	newline = []byte{'\n'} // Byte array for newline
	space   = []byte{' '}  // Byte array for space
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	hub  *Hub            // Reference to the hub managing all clients
	conn *websocket.Conn // WebSocket connection for the client
	send chan []byte     // Channel for sending messages to the client
	done chan struct{}   // Channel for signaling that the client is done (for graceful shutdown)
}

func (c *Client) ReadPump(ctx *gin.Context) {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
		close(c.done)
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	monitorCtx, cancelMonitor := context.WithCancel(context.Background())
	defer cancelMonitor()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		var monitoringData types.MonitorStockPrice
		if err := json.Unmarshal(message, &monitoringData); err != nil {
			log.Printf("Invalid message: %v", err)
			continue
		}

		// Register the client to clientsMap
		c.hub.mu.Lock()
		c.hub.clientsMap[monitoringData.AlertID] = c

		// Add the client to activeTickersMap (allow multiple clients for each ticker)
		if _, exists := c.hub.activeTickersMap[monitoringData.TickerToMonitor]; !exists {
			c.hub.activeTickersMap[monitoringData.TickerToMonitor] = []*Client{c} // Start new list with current client
			// Start monitoring in a separate goroutine
			log.Println("Before starting goroutine, Total Goroutines:", runtime.NumGoroutine())

			go func() {
				log.Println("Goroutine started for AlertID:", monitoringData.AlertID)
				c.monitorStockPrice(monitoringData, monitoringData.AlertID, monitorCtx)
				log.Println("Goroutine exited for AlertID:", monitoringData.AlertID)
			}()

		} else {
			c.hub.activeTickersMap[monitoringData.TickerToMonitor] = append(c.hub.activeTickersMap[monitoringData.TickerToMonitor], c)
		}
		log.Println("After starting goroutine, Total Goroutines:", runtime.NumGoroutine())
		c.hub.mu.Unlock()
	}
}

func (c *Client) monitorStockPrice(monitoringData types.MonitorStockPrice, alertID string, monitorCtx context.Context) {
	var stockData types.StockData
	tickerChan := time.NewTicker(2 * time.Second)
	defer tickerChan.Stop()

	for {
		select {
		case <-monitorCtx.Done():
			// Stop monitoring if no longer needed
			c.hub.mu.Lock()
			clients := c.hub.activeTickersMap[monitoringData.TickerToMonitor]
			for i, client := range clients {
				if client == c {
					c.hub.activeTickersMap[monitoringData.TickerToMonitor] = append(clients[:i], clients[i+1:]...)
					break
				}
			}
			c.hub.mu.Unlock()
			return
		case <-tickerChan.C:
			// Fetch the latest stock price
			currentPrice, currentTime, err := utils.GetCurrentStockPriceAndTime(monitoringData.Ticker, stockData)
			if err != nil {
				log.Printf("Error fetching stock price: %v", err)
				continue
			}

			// Prepare the response
			response := map[string]interface{}{
				"statusCode": http.StatusOK,
				"message":    "Latest price fetched successfully",
				"data": types.GetCurrentPrice{
					CurrentFetchedPrice: currentPrice,
					CurrentFetchedTime:  currentTime,
					AlertID:             alertID,
				},
				"error": nil,
			}

			// Send stock price updates to all clients monitoring the ticker
			responseJSON, err := json.Marshal(response)
			if err != nil {
				log.Printf("Error marshaling response: %v", err)
				continue
			}

			c.hub.mu.RLock()
			for _, client := range c.hub.activeTickersMap[monitoringData.TickerToMonitor] {
				select {
				case client.send <- responseJSON:
				case <-time.After(writeWait):
					log.Printf("Failed to send message to client: timeout")
				}
			}
			c.hub.mu.RUnlock()
		}
	}
}

func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			err := c.conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Printf("Error writing message: %v", err)
				return
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Printf("Error writing ping: %v", err)
				return
			}
		}
	}
}

// UnregisterClientByAlertID will be used to remove a client based on its alert ID
// func (h *Hub) UnregisterClientByAlertID(alertID string) {
// Close the WebSocket connection
// }

func ServeWs(c *gin.Context, hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{
		hub:  hub,
		conn: conn,
		send: make(chan []byte, 256),
		done: make(chan struct{}),
	}

	client.hub.register <- client

	go client.WritePump()
	go client.ReadPump(c)
}
