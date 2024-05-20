package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dr4g0n7ly/AutoTariff-Service/types"
	"github.com/gorilla/websocket"
)

func main() {
	recv, err := NewDataReceiver()
	if err != nil {
		print("fucked")
		log.Fatal(err)
	}
	fmt.Println("starting data receiver")
	http.HandleFunc("/ws", recv.handleWS)
	http.ListenAndServe(":30000", nil)
}

type DataReceiver struct {
	msgch chan OBUData
	conn  *websocket.Conn
	prod  DataProducer
}

type OBUData struct {
	OBUID int     `json:"obuID"`
	Lat   float64 `json:"lat"`
	Long  float64 `json:"long"`
}

func NewDataReceiver() (*DataReceiver, error) {
	kafkaTopic := "obudata"
	p, err := NewKafkaProducer(kafkaTopic)
	if err != nil {
		print("ERROR CREATING NEW PRODUCER")
		return nil, err
	}
	p = NewLogMiddleware(p)
	return &DataReceiver{
		msgch: make(chan OBUData, 128),
		prod:  p,
	}, nil
}

func (dr *DataReceiver) produceData(data types.OBUData) error {
	return dr.prod.ProduceData(data)
}

func (dr *DataReceiver) handleWS(w http.ResponseWriter, r *http.Request) {
	u := websocket.Upgrader{
		ReadBufferSize:  1028,
		WriteBufferSize: 1028,
	}
	conn, err := u.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	dr.conn = conn

	go dr.wsReceiveLoop()
}

func (dr *DataReceiver) wsReceiveLoop() {
	fmt.Println("New OBU connected client connected")

	for {
		var data types.OBUData
		if err := dr.conn.ReadJSON(&data); err != nil {
			log.Println("read error:", err)
			continue
		}
		fmt.Println("received message", data)
		if err := dr.produceData(data); err != nil {
			fmt.Println("kafka producer error:", err)
		}
	}
}
