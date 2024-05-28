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

	// httpClient := client.NewHTTPClient(aggregatorEndpoint)
	grpcClient, err := client.NewGRPCClient(aggregatorEndpoint)
	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, calcServ, grpcClient)
	if err != nil {
		log.Fatal(err)
	}
	kafkaConsumer.Start()
}
