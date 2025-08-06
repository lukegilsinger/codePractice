package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func StartConsumer(broker, topic, groupID string) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{broker},
		Topic:    topic,
		GroupID:  groupID,
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})

	go func() {
		log.Println("Kafka consumer started...")
		for {
			msg, err := reader.ReadMessage(context.Background())
			if err != nil {
				log.Printf("Error reading message: %v", err)
				continue
			}
			log.Printf("Consumed: key=%s value=%s", string(msg.Key), string(msg.Value))
		}
	}()
}
