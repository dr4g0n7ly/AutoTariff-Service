package main

import (
	"fmt"

	"github.com/dr4g0n7ly/AutoTariff-Service/types"
)

type CalculatorServicer interface {
	CalculateDistance(types.OBUData) (float64, error)
}

type CalculatorService struct{}

func NewCalculatorService() CalculatorServicer {
	return &CalculatorService{}
}

func (s *CalculatorService) CalculateDistance(data types.OBUData) (float64, error) {
	fmt.Println("dslkfjasl")
	return 0.0, nil
}
