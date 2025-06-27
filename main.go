package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"restaurantnotificationservice/db"
	"time"
)

type OrderMessage struct {
	OrderId  string `json:"orderId"`
	Username string `json:"foodName"`
	Amount   int    `json:"amount"`
	Money    int    `json:"money"`
	Message  string `json:"message"`
}

func main() {
	db.ConnectionDB()
	// Tạo Kafka Reader (consumer)
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"localhost:9092"}, // Địa chỉ Kafka broker
		GroupID:     "notification-group",       // GroupID cho consumer group (có thể đặt bất kỳ)
		Topic:       "order-events",             // Topic muốn lắng nghe
		MinBytes:    10e2,                       // 1KB
		MaxBytes:    10e6,                       // 10MB
		StartOffset: kafka.FirstOffset,          // Đọc từ đầu topic (nếu cần)
	})
	defer r.Close()

	ctx := context.Background()
	fmt.Println("Start listening for order-events...")

	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			log.Printf("Read error: %v", err)
			time.Sleep(time.Second)
			continue
		}

		// Parse event
		var evt OrderMessage
		if err := json.Unmarshal(m.Value, &evt); err != nil {
			log.Printf("Invalid message: %s", string(m.Value))
			continue
		}
		// Log ra thông tin đơn hàng
		log.Printf("Received order: %+v", evt)
		// Bạn có thể xử lý gửi notification ở đây!
	}
}
