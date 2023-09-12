package main

import (
	"log"

	"github.com/dqwei1219/toll-calculator-project/aggregator/client"
)

const kafkaTopic = "gpu-coordinate"

func main() {
	var (
		svc CalculatorServicer
		err error
	)

	svc, err = NewCalculatorService()
	svc = NewLogMiddleware(svc)

	if err != nil {
		log.Fatal(err) // not going to work if svc is nil
	}
	c, err := NewKafkaComsumer(kafkaTopic, svc, client.NewClient("http://localhost:3000/aggregate"))
	if err != nil {
		log.Fatal(err)
	}
	c.Start()
}
