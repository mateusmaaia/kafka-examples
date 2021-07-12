package main

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	// mechanism := plain.Mechanism{Username: "admin", Password: "admin-secret"}

	// dialer := &kafka.Dialer{
	// 	Timeout:       10 * time.Second,
	// 	DualStack:     true,
	// 	SASLMechanism: mechanism,
	// }

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:29092"},
		Topic:    "topic-name",
		Balancer: &kafka.Hash{},
		// Dialer:   dialer,
	})

	err := w.WriteMessages(
		context.Background(),
		kafka.Message{Value: []byte("testing one!")},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
