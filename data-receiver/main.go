package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dqwei1219/toll-calculator-project/types"
	"github.com/gorilla/websocket"
)


type DataReceiver struct {
	msgch chan types.UnitCoordinate
	conn  *websocket.Conn
	prod  DataProducer
}

func NewDataReceiver() (*DataReceiver, error) {
	var (
		p  DataProducer
		err error
		kafkaTopic = "gpu-coordinate"
	)

	p, err = NewKafkaProducer(kafkaTopic)
	if err != nil {
		return nil, err
	}

	p = NewLogMiddleware(p) // add logging middleware


	dr := &DataReceiver{
		msgch: make(chan types.UnitCoordinate, 128),
		prod:  p,
	}

	return dr, nil
}

func (dr *DataReceiver) wsHandler(w http.ResponseWriter, r *http.Request) {
	u := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, err := u.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	dr.conn = conn
	go dr.receiveData()
}

func (dr *DataReceiver) receiveData() {
	fmt.Println("Start receiving data")
	for {
		var data types.UnitCoordinate
		if err := dr.conn.ReadJSON(&data); err != nil {
			log.Println(err)
			continue
		}
		if err := dr.produceData(data); err != nil {
			fmt.Println("Kafka produce error: ", err)
		}
	}
}

func (dr *DataReceiver) produceData(data types.UnitCoordinate) error {
	return dr.prod.ProduceData(data)
}

func main() {
	/** NewDataReceiver ->
	 *  wsHandler (will create websocker connection, and a kafka producer) ->
	 *	receiveData (will receive data from websocket
									 connection, and produce data to kafka) ->
	 *  produceData (will produce data to kafka)
	*/
	recv, err := NewDataReceiver()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/ws", recv.wsHandler)
	fmt.Println("Server started")
	http.ListenAndServe(":30000", nil)

}
