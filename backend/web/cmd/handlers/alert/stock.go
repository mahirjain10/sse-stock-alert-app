package alert

// import (
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/websocket"
// )

// func StartMonitoringStockTicker() {

// 	// Dial the WebSocket server
// 	conn, _, err := websocket.DefaultDialer.Dial(serverURL, http.Header{})
// 	if err != nil {
// 		log.Fatalf("Failed to connect to WebSocket server: %v", err)
// 	}
// 	defer conn.Close()

// 	err = conn.WriteMessage(websocket.TextMessage, requestData)
// 	if err != nil {
// 		log.Fatalf("Failed to send message: %v", err)
// 	}
// }