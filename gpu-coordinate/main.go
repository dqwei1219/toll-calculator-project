package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
	"github.com/gorilla/websocket"
	"github.com/pingcap/log"
	"github.com/dqwei1219/toll-calculator-project/types"
)

const wsEndpoint = "ws://localhost:30000/ws"
var sendInterval = time.Second * 5


func generateLocation() (float64, float64) {
	return generateCoord(), generateCoord()
}
func generateCoord() float64 {
	// generate random coordinate for mimicking the coordinate from the unit
	n := rand.Intn(100) + 1
	f := rand.Float64()

	return float64(n) + f
}
func generateUnitId(n int) []int {
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = rand.Intn(math.MaxInt32)
	}
	return ids
}

func main() {
	// generate random coordinate for mimicking the coordinate from the unit
	unitIds := generateUnitId(20)
	conn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)
	if (err != nil) {
		log.Error(err.Error())
		return
	}
	for {
		for i := 0; i < len(unitIds); i++ {
			lat, long := generateLocation()
			data := types.UnitCoordinate{
				UnitId: unitIds[i],
				Latitude: lat,
				Longitude: long,
			}
			if err := conn.WriteJSON(data); err != nil {
				log.Error(err.Error())
			}

			fmt.Printf("%+v\n", data)
			time.Sleep(sendInterval)
		}
	}
}


