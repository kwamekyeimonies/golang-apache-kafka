package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"

	"github.com/kwamekyeimonies/Go-Apache-Kafka/config"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error: %v", err.Error())
	}
	fmt.Println("KAFKASERVER1:", os.Getenv("KAFKASERVER1"))
	fmt.Println("CLIENTID:", os.Getenv("CLIENTID"))

	topic := "teaCodeTopix"

	for {
		producer, err := config.ProducerConfig()
		if err != nil {
			log.Fatalf("error: %v", err.Error())
		}
		deliveryChannel := make(chan kafka.Event, 10000)

		err = producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte("Daniel Tenkorang")},
			deliveryChannel,
		)
		if err != nil {
			log.Fatalf("error: %v", err.Error())
		}

		<-deliveryChannel

		time.Sleep(time.Second * 5)
	}

}
