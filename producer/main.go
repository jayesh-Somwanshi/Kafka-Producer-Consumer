package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	topic         = "my-topic"
	brokerAddress = "localhost:9092"
)

// Define a message struct
type Message struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Time    string `json:"time"`
}

func main() {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{brokerAddress},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	defer writer.Close()

	scanner := bufio.NewScanner(os.Stdin)
	id := 1

	for {
		fmt.Print("Enter message: ")
		scanner.Scan()
		content := scanner.Text()

		if content == "" {
			fmt.Println("Empty message, exiting...")
			break
		}

		// Create a message struct
		msg := Message{
			ID:      id,
			Content: content,
			Time:    time.Now().Format(time.RFC3339),
		}
		id++

		// Convert to JSON
		msgBytes, err := json.Marshal(msg)
		if err != nil {
			log.Fatalf("Error encoding message: %v", err)
		}

		err = writer.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte(fmt.Sprintf("Key-%d", id)),
				Value: msgBytes,
			},
		)

		if err != nil {
			log.Fatalf("Failed to send message: %v", err)
		}

		fmt.Println("Message sent:", string(msgBytes))
	}
}
