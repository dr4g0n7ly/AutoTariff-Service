package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	fmt.Println("Checking Kafka using confluent-kafka-go...")
	config := &kafka.ConfigMap{"bootstrap.servers": "localhost:9092"} // Replace with your actual server address
	producer, err := kafka.NewProducer(config)
	if err != nil {
		fmt.Println("Failed to create producer:", err)
		return
	}
	defer producer.Close()

	topic := "test-topic" // Replace with an existing topic

	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic},
		Value:          []byte{},
	}

	err = producer.Produce(msg, nil)
	if err != nil {
		fmt.Println("Kafka server might be unreachable:", err)
	} else {
		fmt.Println("Kafka server appears to be running (using confluent-kafka-go)!")
	}
}
