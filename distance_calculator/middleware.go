package main

import (
	"time"

	"github.com/dr4g0n7ly/AutoTariff-Service/types"
	"github.com/sirupsen/logrus"
)

type LogMiddleware struct {
	next CalculatorServicer
}

func NewLogMiddleware(next CalculatorServicer) CalculatorServicer {
	return &LogMiddleware{
		next: next,
	}
}

func (m *LogMiddleware) CalculateDistance(data types.OBUData) (distance float64, err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"took": time.Since(start),
			"err":  err,
			"dist": distance,
		}).Info("calculate distance")
	}(time.Now())
	distance, err = m.next.CalculateDistance(data)
	return distance, err
}
