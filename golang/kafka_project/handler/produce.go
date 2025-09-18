package handler

import (
	"encoding/json"
	"net/http"

	"rest-kafka-service/kafka"
	"rest-kafka-service/models"
)

type ProduceHandler struct {
	Producer *kafka.KafkaProducer
}

func NewProduceHandler(producer *kafka.KafkaProducer) *ProduceHandler {
	return &ProduceHandler{Producer: producer}
}

func (h *ProduceHandler) HandleProduce(w http.ResponseWriter, r *http.Request) {
	var msg models.Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	err := h.Producer.ProduceMessage(r.Context(), msg.Key, msg.Value)
	if err != nil {
		http.Error(w, "Failed to send to Kafka", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Message sent to Kafka"))
}
