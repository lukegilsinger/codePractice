package main

import (
	"log"
	"net/http"

	"rest-kafka-service/handler"
	"rest-kafka-service/kafka"
)

const (
	kafkaBroker = "localhost:9092"
	kafkaTopic  = "my-topic"
	groupID     = "my-consumer-group"
)

func main() {
	// Start Kafka consumer in background
	kafka.StartConsumer(kafkaBroker, kafkaTopic, groupID)

	// Create Kafka producer
	producer := kafka.NewKafkaProducer(kafkaBroker, kafkaTopic)

	// Create HTTP handler
	produceHandler := handler.NewProduceHandler(producer)

	// Register routes
	http.HandleFunc("/produce", produceHandler.HandleProduce)

	log.Println("REST server listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
