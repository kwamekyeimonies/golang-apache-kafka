package config

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ProducerConfig() (*kafka.Producer, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKASERVER1"),
		"client.id":         os.Getenv("CLIENTID"),
		"acks":              "all",
	})

	if err != nil {
		return nil, fmt.Errorf("Failed to create producer: %v\n", err.Error())
	}

	return producer, nil
}
