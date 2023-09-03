package main

import (
	"fmt"
	"math/rand"
	"time"
)

var sendInterval = time.Second

type unitCoordinate struct {
	unitId 		int 			`json:"unitId"`
	latitude 	float64 	`json:"latitude"`
	longitude float64 	`json:"longitude"`
}

func generateCoord() float64 {
	// generate random coordinate for mimicking the coordinate from the unit
	n := rand.Intn(100) + 1
	f := rand.Float64()

	return float64(n) + f
}

func main() {
	for {
		fmt.Println(generateCoord())
		time.Sleep(sendInterval)
	}
}


