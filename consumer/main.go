// package main

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"log"

// 	"github.com/segmentio/kafka-go"
// )

// const (
// 	topic         = "my-topic"
// 	brokerAddress = "localhost:9092"
// 	groupID       = "my-group"
// )

// // Define message struct
// type Message struct {
// 	ID      int    `json:"id"`
// 	Content string `json:"content"`
// 	Time    string `json:"time"`
// }

// func main() {
// 	reader := kafka.NewReader(kafka.ReaderConfig{
// 		Brokers: []string{brokerAddress},
// 		Topic:   topic,
// 		GroupID: groupID,
// 	})

// 	defer reader.Close()

// 	fmt.Println("Consumer started... Listening for messages.")

// 	for {
// 		msg, err := reader.ReadMessage(context.Background())
// 		if err != nil {
// 			log.Fatalf("Failed to read message: %v", err)
// 		}

// 		// Decode JSON
// 		var receivedMsg Message
// 		err = json.Unmarshal(msg.Value, &receivedMsg)
// 		if err != nil {
// 			log.Printf("Error decoding message: %v", err)
// 			continue
// 		}

// 		fmt.Printf("Received: %+v\n", receivedMsg)
// 	}
// }

package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/segmentio/kafka-go"
)

const (
	topic         = "my-topic"
	brokerAddress = "localhost:9092"
	groupID       = "my-group"
)

// Define message struct
type Message struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Time    string `json:"time"`
}

func main() {
	// Connect to MySQL
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/kafka_messages")
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}
	defer db.Close()

	// Create Kafka reader
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: groupID,
	})

	defer reader.Close()

	fmt.Println("Consumer started... Listening for messages.")

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("Failed to read message: %v", err)
		}

		// Decode JSON
		var receivedMsg Message
		err = json.Unmarshal(msg.Value, &receivedMsg)
		if err != nil {
			log.Printf("Error decoding message: %v", err)
			continue
		}

		// Insert message into MySQL
		_, err = db.Exec("INSERT INTO messages (content, received_time) VALUES (?, ?)", receivedMsg.Content, receivedMsg.Time)
		if err != nil {
			log.Printf("Error inserting into database: %v", err)
			continue
		}

		fmt.Printf("Saved to DB: %+v\n", receivedMsg)
	}
}
