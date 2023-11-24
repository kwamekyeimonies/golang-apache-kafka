package main

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func main() {

	topic := "teaCodeTopix"
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"client.id":         "hello_something",
		"acks":              "all",
	})

	if err != nil {
		fmt.Printf("Failed to create producer: %v\n", err.Error())
	}

	go func() {
		consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
			"bootstrap.servers": "localhost:9092",
			"group.id":          "hello_something",
			"auto.offset.reset": "smallest",
		})

		if err != nil {
			log.Fatal(err.Error())
		}

		err = consumer.Subscribe(topic, nil)
		if err != nil {
			log.Fatal(err)
		}

		for {
			ev := consumer.Poll(100)
			switch e := ev.(type) {
			case *kafka.Message:
				fmt.Printf("Consumed message from queue: %+v\n", string(e.Value))
			case *kafka.Error:
				fmt.Printf("%+v\n", e)
			}
		}
	}()

	deliverchannel := make(chan kafka.Event, 10000)
	for {
		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte("Dexoangle"),
		},
			deliverchannel,
		)

		if err != nil {
			log.Fatal(err)
		}

		<-deliverchannel
		// fmt.Printf("%+v\n", e.String())
		time.Sleep(time.Second * 5)

	}

	// fmt.Printf("%+v\n", p)
}
