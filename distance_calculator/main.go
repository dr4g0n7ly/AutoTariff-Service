package main

import (
	"log"

	"github.com/dr4g0n7ly/AutoTariff-Service/aggregator/client"
)

const kafkaTopic = "obudata"
const aggregatorEndpoint = "http://127.0.0.1:3000/aggregate"

func main() {
	calcServ := NewCalculatorService()
	calcServ = NewLogMiddleware(calcServ)
	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, calcServ, client.NewClient(aggregatorEndpoint))
	if err != nil {
		log.Fatal(err)
	}
	kafkaConsumer.Start()
}
