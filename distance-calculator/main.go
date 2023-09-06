package main

import (
	"log"
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
	c, err := NewKafkaComsumer(kafkaTopic, svc)
	if err != nil {
		log.Fatal(err)
	}
	c.Start()
}