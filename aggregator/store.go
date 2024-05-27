package main

import (
	"fmt"

	"github.com/dr4g0n7ly/AutoTariff-Service/types"
)

type Storer interface {
	Insert(types.Distance) error
	GetDistance(int) (float64, error)
}

type MemoryStore struct {
	data map[int]float64
}

func (m *MemoryStore) Insert(d types.Distance) error {
	m.data[d.OBUID] += d.Distance
	return nil
}

func (m *MemoryStore) GetDistance(obuID int) (float64, error) {
	dist, ok := m.data[obuID]
	if !ok {
		return 0.0, fmt.Errorf("Error retrieving distance")
	}
	return dist, nil
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[int]float64),
	}
}
