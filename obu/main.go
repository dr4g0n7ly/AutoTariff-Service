package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/dr4g0n7ly/AutoTariff-Service/types"
	"github.com/gorilla/websocket"
)

const wsEndpoint = "ws://127.0.0.1:30000/ws"

var sendInterval = time.Second * 5

func genCoord() float64 {
	n := float64(rand.Intn(100) + 1)
	f := rand.Float64()
	return n + f
}

func genLocation() (float64, float64) {
	return genCoord(), genCoord()
}

func genOBUs(n int) []int {
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = rand.Intn(99999)
	}
	return ids
}

func main() {
	rand.Seed(time.Now().UnixNano())
	conn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	obuIDs := genOBUs(20)
	for {
		for i := 0; i < len(obuIDs); i++ {
			data := types.OBUData{
				OBUID: obuIDs[i],
				Lat:   genCoord(),
				Long:  genCoord(),
			}
			fmt.Printf("%+v\n", data)
			if err := conn.WriteJSON(data); err != nil {
				log.Fatal(err)
			}
		}
		time.Sleep(sendInterval)
	}
}
