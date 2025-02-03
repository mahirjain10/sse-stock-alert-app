package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mahirjain_10/stock-alert-app/backend/internal/types"
	"github.com/redis/go-redis/v9"
)

const (
	socketServerURL = "ws://localhost:8080/ws/get-stock-price-socket"
	redisChannel    = "monitor"
)
var count int = 0;

// Publish sends a message to the Redis Pub/Sub channel.
func Publish(redisClient *redis.Client, ctx context.Context, ticker, alertID string) {
	ist, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		log.Println("Error loading IST timezone:", err)
		return
	}
	currentTime := time.Now().In(ist).Format("Jan 02, 2006 at 3:04pm (MST)")

	message := map[string]string{
		"ticker_to_monitor": ticker,
		"alert_id":          alertID,
	}
	messageJSON, err := json.Marshal(message)
	if err != nil {
		log.Printf("[%s] - Error marshaling message: %s\n", currentTime, err)
		return
	}

	if err := redisClient.Publish(ctx, redisChannel, string(messageJSON)).Err(); err != nil {
		log.Printf("[%s] - Error publishing message: %s\n", currentTime, err)
		return
	}

	log.Printf("[%s] - Message published: %s\n", currentTime, string(messageJSON))
}

// Subscribe listens for messages on the Redis Pub/Sub channel.
func Subscribe(redisClient *redis.Client, ctx context.Context) {
	// ctx = context.WithValue(ctx,"count",count)
	ist, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		log.Println("Error loading IST timezone:", err)
		return
	}

	pubsub := redisClient.Subscribe(ctx, redisChannel)
	defer pubsub.Close()
	log.Println("Listening for messages...")

	// retrieveCtx := context.Background()
	var conn *websocket.Conn

	// Connect to WebSocket and manage connection
	connectWebSocket := func()error {
		var dialErr error
		conn, _, dialErr = websocket.DefaultDialer.Dial(socketServerURL, nil)
		if dialErr != nil {
			log.Println("WebSocket dial error:", dialErr)
			return dialErr
		}
		startWebSocketReader(conn, redisClient, ctx, ist)
		return nil
	}

	for {
		if err = connectWebSocket(); err == nil {
			break
		}
		log.Println("Retrying WebSocket connection in 5 seconds...")
		time.Sleep(5 * time.Second)
	}
	defer conn.Close()

	processRedisMessages(pubsub, conn, connectWebSocket, ctx, ist)
}

// READ MESSAGE AND MONITORS THE PRICES USING FUNC
func startWebSocketReader(conn *websocket.Conn, redisClient *redis.Client, ctx context.Context, ist *time.Location) {
	// ctxWithTimeOut,cancel() := context.WithCancel(ctx)

	go func() {
		for {
			_, response, err := conn.ReadMessage()
			if err != nil {
				log.Printf("WebSocket read error: %v", err)
				return
			}

			currentTime := time.Now().In(ist).Format("Jan 02, 2006 at 3:04pm (MST)")
			log.Printf("[%s] - WebSocket response: %s\n", currentTime, response)

			var parsedResponse types.Response
			if err := json.Unmarshal(response, &parsedResponse); err != nil {
				log.Printf("Error unmarshaling response: %v", err)
				continue
			}
			if parsedResponse.Data == nil {
				return
			}

			dataJSON, err := json.Marshal(parsedResponse.Data)
			if err != nil {
				log.Printf("Error marshaling response data: %v", err)
				return
			}

			var stockData types.GetCurrentPrice
			if err := json.Unmarshal(dataJSON, &stockData); err != nil {
				log.Printf("Error unmarshaling stock data: %v", err)
				return
			}
			fmt.Println("I am getting called in startwebsocket func")
			go ComparePriceAndThreshold(redisClient, ctx, stockData.AlertID, int64(stockData.CurrentFetchedPrice))
		}
	}()
}

func ComparePriceAndThreshold(redisClient *redis.Client, ctx context.Context, alertID string, currentPrice int64) {
	count=count+1;
	alertData, err := redisClient.HGetAll(ctx, alertID).Result()
	if err != nil {
		log.Printf("Error retrieving alert data from Redis: %v", err)
		return
	}

	alertPrice, err := strconv.ParseInt(alertData["alert_price"], 10, 64)
	if err != nil {
		log.Printf("Error parsing alert price: %v", err)
		return
	}
	fmt.Println(count)
	if count == 10 {
		currentPrice = alertPrice
		fmt.Println("current price := ",currentPrice)
	}
	isConditionMet, err := CompareUsingSymbol(alertData["alert_condition"], currentPrice, alertPrice)
	if err != nil {
		log.Printf("Error evaluating alert condition: %v", err)
		return
	}
	responseData := types.UpdateActiveStatus{
		UserID: alertData["user_id"],
		ID:     alertData["alert_id"],
		Active: false,
	}
	if isConditionMet {
		log.Printf("Alert condition met for alert ID: %s\n", alertID)
		err := PublishToPubSub(redisClient, ctx, "alert-topic", responseData)
		fmt.Println("calling pub sub ")
		if err != nil {
			log.Printf("Error publishing to Pub/Sub: %v", err)
		}
	} else {
		log.Printf("Alert condition not met for alert ID: %s\n", alertID)
	}
}

// PROCESSES THE INCOMING MESSAGE FROM PUBLISH AND SENDS THE MESSAGE TO WEBSOCKET
func processRedisMessages(pubsub *redis.PubSub, conn *websocket.Conn, reconnectFunc func() error, ctx context.Context, ist *time.Location) {
	for {
		msg, err := pubsub.ReceiveMessage(ctx)
		if err != nil {
			log.Printf("Error receiving Redis message: %s\n", err)
			continue
		}

		currentTime := time.Now().In(ist).Format("Jan 02, 2006 at 3:04pm (MST)")
		log.Printf("[%s] - Received message: %s\n", currentTime, msg.Payload)

		if err := conn.WriteMessage(websocket.TextMessage, []byte(msg.Payload)); err != nil {
			log.Println("WebSocket write error. Reconnecting...")
			conn.Close()

			for {
				if err = reconnectFunc(); err == nil {
					break
				}
				log.Println("Retrying WebSocket connection in 5 seconds...")
				time.Sleep(5 * time.Second)
			}
		}
	}
}
