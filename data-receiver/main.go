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
	conn *websocket.Conn
}

func (dr *DataReceiver) wsHandler (w http.ResponseWriter, r *http.Request)  {
	u := websocket.Upgrader{
		ReadBufferSize: 1024,
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
			continue;
		}
		fmt.Printf("received data from [%d] : <lat %.2f, long %.2f>\n",
		 					  data.UnitId, data.Latitude, data.Longitude,
							)
		dr.msgch <- data
	}
}

func NewDataReceiver() *DataReceiver {
	dr := &DataReceiver{
		msgch: make(chan types.UnitCoordinate, 128),
	}
	return dr
}

func main() {
	recv := NewDataReceiver()
	http.HandleFunc("/ws", recv.wsHandler)
	fmt.Println("Server started")
	http.ListenAndServe(":30000", nil)
}
