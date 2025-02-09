package test

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	serverURL   = "ws://localhost:8080/ws/get-stock-price-socket" // Replace with your WebSocket server URL
	connections = 1000  // Number of concurrent WebSocket connections to test
	maxRetries  = 3     // Number of times to retry on failure
)

// Function to connect and send messages
func connectAndSend(wg *sync.WaitGroup, id int) {
	defer wg.Done()

	headers := http.Header{}
	var conn *websocket.Conn
	var err error

	// Retry mechanism for failed connections
	for retry := 0; retry < maxRetries; retry++ {
		conn, _, err = websocket.DefaultDialer.Dial(serverURL, headers)
		if err == nil {
			break
		}
		log.Printf("Connection %d failed (retry %d/%d): %v", id, retry+1, maxRetries, err)
		time.Sleep(2 * time.Second) // Wait before retrying
	}

	if err != nil {
		log.Printf("Connection %d failed after %d retries: %v", id, maxRetries, err)
		return
	}
	defer conn.Close()

	// Send a subscription message
	err = conn.WriteMessage(websocket.TextMessage, []byte(`{"ticker_to_monitor":"LICI.NS","alert_id":"2u389y4734ekjs"}`))
	if err != nil {
		log.Printf("Error sending message on connection %d: %v", id, err)
		return
	}

	// Listen for messages
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Connection %d closed: %v", id, err)
			return
		}
		fmt.Printf("Connection %d received: %s\n", id, msg)
		time.Sleep(500 * time.Millisecond) // Adjust processing delay
	}
}

// Load testing function
func Wstest() {
	var wg sync.WaitGroup

	// Launch multiple WebSocket connections
	for i := 0; i < connections; i++ {
		wg.Add(1)
		go connectAndSend(&wg, i)

		// Reduce delay to stress test more effectively
		if i%100 == 0 {
			time.Sleep(100 * time.Millisecond)
		}
	}

	wg.Wait()
}
