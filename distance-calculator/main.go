package main

import (
	"log"

	"github.com/dqwei1219/toll-calculator-project/aggregator/client"
)

const kafkaTopic = "gpu-coordinate"
const Endpoint = "http://localhost:3000/aggregate"

func main() {
	var (
		svc CalculatorServicer
		err error
	)

	svc, err = NewCalculatorService()
	svc = NewLogMiddleware(svc)

	// httpClient := client.NewHTTPClient(Endpoint)
	grpcClient, err := client.NewGRPCClient(Endpoint)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err) // not going to work if svc is nil
	}
	c, err := NewKafkaComsumer(kafkaTopic, svc, grpcClient)
	if err != nil {
		log.Fatal(err)
	}
	c.Start()
}
