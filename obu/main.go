package main

import (
	"fmt"
	"math"
	"math/rand/v2"
	"time"
)

var sendInterval = time.Second

type OBUData struct {
	OBUID int     `json:"obuID"`
	Lat   float64 `json:"lat"`
	Long  float64 `json:"long"`
}

func genCoord() float64 {
	n := float64(rand.IntN(100) + 1)
	f := rand.Float64()
	return n + f
}

func genLocation() (float64, float64) {
	return genCoord(), genCoord()
}

func genOBUs(n int) []int {
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = rand.IntN(math.MaxInt)
	}
	return ids
}

func main() {
	obuIDs := genOBUs(20)
	for {
		for i := 0; i < len(obuIDs); i++ {
			data := OBUData{
				OBUID: obuIDs[i],
				Lat:   genCoord(),
				Long:  genCoord(),
			}
			fmt.Printf("%+v\n", data)
		}
		time.Sleep(sendInterval)
	}
}
