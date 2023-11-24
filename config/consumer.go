package config

import (
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ConsumerConfig() (*kafka.Consumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKASERVER1"),
		"group.id":          os.Getenv("CONSUMERGROUPID"),
		"auto.offset.reset": "smallest",
	})

	if err != nil {
		return nil, err
	}

	return consumer, nil
}
