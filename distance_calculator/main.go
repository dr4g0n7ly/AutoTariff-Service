package main

import (
	"log"
)

const kafkaTopic = "obudata"

func main() {
	calcServ := NewCalculatorService()
	calcServ = NewLogMiddleware(calcServ)
	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, calcServ)
	if err != nil {
		log.Fatal(err)
	}
	kafkaConsumer.Start()
}
