package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	Writer *kafka.Writer
}

func NewKafkaProducer(broker, topic string) *KafkaProducer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker},
		Topic:   topic,
	})

	return &KafkaProducer{Writer: writer}
}

func (kp *KafkaProducer) ProduceMessage(ctx context.Context, key, value string) error {
	err := kp.Writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(key),
		Value: []byte(value),
	})
	if err != nil {
		log.Printf("Failed to write message: %v", err)
	}
	return err
}
